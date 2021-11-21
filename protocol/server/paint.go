package server

import "Projector/protocol/common"

type WindowPaintEvent interface {
	WindowEvent

	windowPaintEvent()
}

// PaintArcEvent @SerialName("a")
type PaintArcEvent struct {
	WindowPaintEvent

	PaintType  common.PaintType `json:"a"`
	X          int              `json:"b"`
	Y          int              `json:"c"`
	Width      int              `json:"d"`
	Height     int              `json:"e"`
	StartAngle int              `json:"f"`
	ArcAngle   int              `json:"g"`
}

// PaintOvalEvent @SerialName("b")
type PaintOvalEvent struct {
	WindowPaintEvent

	PaintType common.PaintType `json:"a"`
	X         int              `json:"b"`
	Y         int              `json:"c"`
	Width     int              `json:"d"`
	Height    int              `json:"e"`
}

// PaintRoundRectEvent @SerialName("c")
type PaintRoundRectEvent struct {
	WindowPaintEvent

	PaintType common.PaintType `json:"a"`
	X         int              `json:"b"`
	Y         int              `json:"c"`
	Width     int              `json:"d"`
	Height    int              `json:"e"`
	ArcWidth  int              `json:"f"`
	ArcHeight int              `json:"g"`
}

// PaintRectEvent @SerialName("d")
type PaintRectEvent struct {
	WindowPaintEvent

	PaintType common.PaintType `json:"a"`
	X         float64          `json:"b"`
	Y         float64          `json:"c"`
	Width     float64          `json:"d"`
	Height    float64          `json:"e"`
}

// DrawLineEvent @SerialName("e")
type DrawLineEvent struct {
	WindowPaintEvent

	X1 int `json:"a"`
	Y1 int `json:"b"`
	X2 int `json:"c"`
	Y2 int `json:"d"`
}

// CopyAreaEvent @SerialName("f")
type CopyAreaEvent struct {
	WindowPaintEvent

	X      int `json:"a"`
	Y      int `json:"b"`
	Width  int `json:"c"`
	Height int `json:"d"`
	Dx     int `json:"e"`
	Dy     int `json:"f"`
}

// DrawImageEvent @SerialName("l")
type DrawImageEvent struct {
	WindowPaintEvent

	ImageId        common.ImageId        `json:"a"`
	ImageEventInfo common.ImageEventInfo `json:"b"`
}

// DrawStringEvent @SerialName("m")
type DrawStringEvent struct {
	WindowPaintEvent

	Str          string  `json:"a"`
	X            float64 `json:"b"`
	Y            float64 `json:"c"`
	DesiredWidth float64 `json:"d"`
}

// PaintPolygonEvent @SerialName("n")
type PaintPolygonEvent struct {
	WindowPaintEvent

	PaintType common.PaintType `json:"a"`
	Points    []common.Point   `json:"b"`
}

// DrawPolylineEvent @SerialName("o")
type DrawPolylineEvent struct {
	WindowPaintEvent

	Points []common.Point `json:"a"`
}

// PaintPathEvent @SerialName("q")
type PaintPathEvent struct {
	WindowPaintEvent

	PaintType common.PaintType  `json:"a"`
	Path      common.CommonPath `json:"b"`
}
