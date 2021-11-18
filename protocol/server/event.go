package server

import "Projector/protocol/common"

type Event interface {
	event()
}

type ImageDataReplyEvent struct {
	Event
	ImageId   common.ImageId   `json:"a"`
	ImageData common.ImageData `json:"b"`
}

func ToEvent(t string, c map[string]interface{}) Event {
	switch t {
	case "a":
		return &ImageDataReplyEvent{
			ImageId:   common.ToImageId(c["a"].([]interface{})),
			ImageData: common.ToImageData(c["b"].([]interface{})),
		}
	case "e":
		return &DrawCommandsEvent{
			Target:     ToTarget(c["a"].([]interface{})),
			DrawEvents: ToWindowEvents(c["b"].([]interface{})),
		}
	default:
		return nil
	}
}
