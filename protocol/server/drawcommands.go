package server

// DrawCommandsEvent @SerialName("e")
type DrawCommandsEvent struct {
	Event
	Target     DrawCommandsTarget `json:"a"`
	DrawEvents []WindowEvent      `json:"b"`
}

type DrawCommandsTarget interface {
	target()
}

// DrawCommandsTargetOnscreen @SerialName("a")
type DrawCommandsTargetOnscreen struct {
	DrawCommandsTarget

	WindowId int `json:"a"`
}

// DrawCommandsTargetOffscreen @SerialName("b")
type DrawCommandsTargetOffscreen struct {
	DrawCommandsTarget

	PVolatileImageId int64 `json:"a"`
	Width            int   `json:"b"`
	Height           int   `json:"c"`
}

func ToTarget(i []interface{}) DrawCommandsTarget {
	t := i[0].(string)
	c := i[1].(map[string]interface{})

	switch t {
	case "a":
		return &DrawCommandsTargetOnscreen{
			WindowId: int(c["a"].(float64)),
		}
	case "b":
		return &DrawCommandsTargetOffscreen{
			PVolatileImageId: int64(c["a"].(float64)),
			Width:            int(c["b"].(float64)),
			Height:           int(c["c"].(float64)),
		}
	default:
		panic("invalid draw command target type")
	}
}
