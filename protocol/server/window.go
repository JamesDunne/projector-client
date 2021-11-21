package server

import "Projector/protocol/common"
import . "Projector/protocol/util"

type WindowEvent interface {
	windowEvent()
}

func ToWindowEvents(a []interface{}) (events []WindowEvent) {
	events = make([]WindowEvent, len(a))
	for i, o := range a {
		e := o.([]interface{})
		events[i] = ToWindowEvent(e)
	}
	return
}

func ToWindowEvent(i []interface{}) WindowEvent {
	t := i[0].(string)
	c := i[1].(map[string]interface{})

	_ = c
	switch t {
	case "a": // ServerPaintArcEvent
		return &PaintArcEvent{
			PaintType:  common.ToPaintType(c["a"].(string)),
			X:          Jint(c["b"]),
			Y:          Jint(c["c"]),
			Width:      Jint(c["d"]),
			Height:     Jint(c["e"]),
			StartAngle: Jint(c["f"]),
			ArcAngle:   Jint(c["g"]),
		}
	case "b": // ServerPaintOvalEvent
		return &PaintOvalEvent{
			PaintType: common.ToPaintType(c["a"].(string)),
			X:         Jint(c["b"]),
			Y:         Jint(c["c"]),
			Width:     Jint(c["d"]),
			Height:    Jint(c["e"]),
		}
	case "c": // ServerPaintRoundRectEvent
		return &PaintRoundRectEvent{
			PaintType: common.ToPaintType(c["a"].(string)),
			X:         Jint(c["b"]),
			Y:         Jint(c["c"]),
			Width:     Jint(c["d"]),
			Height:    Jint(c["e"]),
			ArcWidth:  Jint(c["f"]),
			ArcHeight: Jint(c["g"]),
		}
	case "d": // ServerPaintRectEvent
		return &PaintRectEvent{
			PaintType: common.ToPaintType(c["a"].(string)),
			X:         c["b"].(float64),
			Y:         c["c"].(float64),
			Width:     c["d"].(float64),
			Height:    c["e"].(float64),
		}
	case "e": // ServerDrawLineEvent
		return &DrawLineEvent{
			X1: Jint(c["a"]),
			Y1: Jint(c["b"]),
			X2: Jint(c["c"]),
			Y2: Jint(c["d"]),
		}
	case "f": // ServerCopyAreaEvent
		return &CopyAreaEvent{
			X:      Jint(c["a"]),
			Y:      Jint(c["b"]),
			Width:  Jint(c["c"]),
			Height: Jint(c["d"]),
			Dx:     Jint(c["e"]),
			Dy:     Jint(c["f"]),
		}
	case "g": // ServerSetFontEvent
		return nil
	case "h": // ServerSetClipEvent
		return nil
	case "i": // ServerSetStrokeEvent
		return nil
	case "j": // ServerDrawRenderedImageEvent
		return nil
	case "k": // ServerDrawRenderableImageEvent
		return nil
	case "l": // ServerDrawImageEvent
		return DrawImageEvent{
			ImageId:        common.ToImageId(c["a"].([]interface{})),
			ImageEventInfo: common.ToImageEventInfo(c["b"].([]interface{})),
		}
	case "m": // ServerDrawStringEvent
		return nil
	case "n": // ServerPaintPolygonEvent
		return nil
	case "o": // ServerDrawPolylineEvent
		return nil
	case "p": // ServerSetTransformEvent
		return nil
	case "q": // ServerPaintPathEvent
		return nil
	case "r": // ServerSetCompositeEvent
		return nil
	case "s": // ServerSetPaintEvent
		return nil
	case "t": // ServerSetUnknownStrokeEvent
		return nil
	default:
		panic("invalid window event type")
	}
}
