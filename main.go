package main

import (
	"Projector/protocol/server"
	"bufio"
	"context"
	"encoding/json"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
	"log"
	"net"
	"os"
	"sync"
)

func main() {
	app.Main(appMain)
	log.Println("main exited")
}

type ClientToServerHandshake struct {
	CommonVersion                 int                   `json:"commonVersion"`
	CommonVersionId               int                   `json:"commonVersionId"`
	Token                         *string               `json:"token"`
	ClientDoesWindowManagement    bool                  `json:"clientDoesWindowManagement"`
	Displays                      []*DisplayDescription `json:"displays"`
	SupportedToClientCompressions []CompressionType     `json:"supportedToClientCompressions"`
	SupportedToClientProtocols    []ProtocolType        `json:"supportedToClientProtocols"`
	SupportedToServerCompressions []CompressionType     `json:"supportedToServerCompressions"`
	SupportedToServerProtocols    []ProtocolType        `json:"supportedToServerProtocols"`
}

type DisplayDescription struct {
	X           int     `json:"x"`
	Y           int     `json:"y"`
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	ScaleFactor float64 `json:"scaleFactor"`
}

type CompressionType string

const (
	NONE CompressionType = "NONE"
	GZIP CompressionType = "GZIP"
)

type ProtocolType string

const (
	KOTLINX_JSON     ProtocolType = "KOTLINX_JSON"
	KOTLINX_PROTOBUF ProtocolType = "KOTLINX_PROTOBUF"
)

func appMain(a app.App) {
	urlstr := "ws://127.0.0.1:9999/"
	args := os.Args[1:]
	if len(args) > 0 {
		urlstr = args[0]
	}

	var glctx gl.Context
	var sz size.Event

	var wsConnectOnce sync.Once

	events := make(chan interface{})
	wsctx, wscancel := context.WithCancel(context.Background())

eventLoop:
	for {
		select {
		// handle GUI event:
		case e := <-a.Events():
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				log.Println(e)
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ = e.DrawContext.(gl.Context)
					onGLInitialize(glctx)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					onGLShutdown(glctx)
					glctx = nil
				}
				if e.To == lifecycle.StageDead {
					break eventLoop
				}
			case size.Event:
				sz = e
				log.Println(sz)

				// connect to websocket once we have window size info:
				wsConnectOnce.Do(func() {
					go connectWebsocket(wsctx, events, urlstr, sz)
				})
			case key.Event:
				log.Println(e)
			case paint.Event:
				if glctx == nil {
					continue
				}
				if e.External {
					continue
				}
				fps.Draw(sz)
				a.Publish()
				//a.Send(paint.Event{External: true})
			}
		case j := <-events:
			if j == nil {
				log.Println("server closed connection")
				break eventLoop
			}
			log.Printf("server: %+v\n", j)

			if arr, ok := j.([]interface{}); ok {
				for _, ia := range arr {
					a, ok := ia.([]interface{})
					if !ok {
						continue
					}

					t := a[0].(string)
					c := a[1].(map[string]interface{})

					e := server.ToEvent(t, c)
					_ = e
					log.Printf("deser: %#v\n", e)
				}
			}
		}
	}

	wscancel()

	log.Println("done")
}

func connectWebsocket(ctx context.Context, events chan interface{}, urlstr string, sz size.Event) {
	var err error

	var wsconn net.Conn
	var wsreader *bufio.Reader
	wsconn, wsreader, _, err = ws.Dial(ctx, urlstr)
	if err != nil {
		log.Println(err)
		return
	}

	_, _ = wsconn, wsreader
	wsw := wsutil.NewWriter(wsconn, ws.StateClientSide, ws.OpText)
	wsr := wsutil.NewClientSideReader(wsconn)
	wsje := json.NewEncoder(wsw)
	wsjd := json.NewDecoder(wsr)

	// write HANDSHAKE_VERSION:
	_, err = wsw.Write([]byte("456250626;0"))
	if err != nil {
		log.Println(err)
		return
	}
	err = wsw.Flush()
	if err != nil {
		log.Println(err)
		return
	}

	// write handshake:
	err = wsje.Encode(&ClientToServerHandshake{
		CommonVersion:              1980644253,
		CommonVersionId:            61 - 46,
		Token:                      nil,
		ClientDoesWindowManagement: false,
		Displays: []*DisplayDescription{
			{
				X:           0,
				Y:           0,
				Width:       sz.WidthPx,
				Height:      sz.HeightPx,
				ScaleFactor: 1.0,
			},
		},
		SupportedToClientCompressions: []CompressionType{NONE},
		SupportedToClientProtocols:    []ProtocolType{KOTLINX_JSON},
		SupportedToServerCompressions: []CompressionType{NONE},
		SupportedToServerProtocols:    []ProtocolType{KOTLINX_JSON},
	})
	if err != nil {
		log.Println(err)
		return
	}
	err = wsw.Flush()
	if err != nil {
		log.Println(err)
		return
	}

	var hdr ws.Header
	var handshakeEvent map[string]interface{}
	var m interface{}

	// expect `type:org.jetbrains.projector.common.protocol.handshake.ToClientHandshakeSuccessEvent`
	hdr, err = wsr.NextFrame()
	if err != nil {
		log.Printf("server: %v\n", err)
		goto done
	}

	_ = hdr
	err = wsjd.Decode(&handshakeEvent)
	if err != nil {
		log.Printf("server: %v\n", err)
		goto done
	}

	log.Printf("server: %+v\n", handshakeEvent)
	if handshakeEvent["type"].(string) != "org.jetbrains.projector.common.protocol.handshake.ToClientHandshakeSuccessEvent" {
		log.Printf("server: handshake failed: reason %v\n", handshakeEvent["reason"].(string))
		goto done
	}

	// TODO:
	//handshakeEvent["colors"]
	//handshakeEvent["fontDataHolders"]

	// finish handshake:
	_, err = wsw.Write([]byte("Unused string meaning fonts loading is done"))
	if err != nil {
		log.Printf("server: %v\n", err)
		goto done
	}
	err = wsw.Flush()
	if err != nil {
		log.Printf("server: %v\n", err)
		goto done
	}

	// receive server events:
	for {
		hdr, err = wsr.NextFrame()
		if err != nil {
			log.Printf("server: %v\n", err)
			break
		}

		err = wsjd.Decode(&m)
		if err != nil {
			log.Printf("server: %v\n", err)
			break
		}

		events <- m
	}

done:
	log.Println("server: close(events)")
	close(events)

	<-ctx.Done()
}

var images *glutil.Images
var fps *debug.FPS

func onGLInitialize(glctx gl.Context) {
	log.Println("GL init")
	images = glutil.NewImages(glctx)
	fps = debug.NewFPS(images)
}

func onGLShutdown(glctx gl.Context) {
	log.Println("GL shutdown")
	fps.Release()
	images.Release()
}
