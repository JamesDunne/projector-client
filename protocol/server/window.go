package server

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
		return nil
	case "b": // ServerPaintOvalEvent
		return nil
	case "c": // ServerPaintRoundRectEvent
		return nil
	case "d": // ServerPaintRectEvent
		return nil
	case "e": // ServerDrawLineEvent
		return nil
	case "f": // ServerCopyAreaEvent
		return nil
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
		return nil
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
