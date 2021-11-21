package common

import . "Projector/protocol/util"

type ImageEventInfo interface {
	imageEventInfo()
}

// ImageEventInfoXy @SerialName("a")
type ImageEventInfoXy struct {
	ImageEventInfo

	X                   int  `json:"a"`
	Y                   int  `json:"b"`
	ArgbBackgroundColor *int `json:"c"`
}

// ImageEventInfoXyWh @SerialName("b")
type ImageEventInfoXyWh struct {
	ImageEventInfo

	X                   int  `json:"a"`
	Y                   int  `json:"b"`
	Width               int  `json:"c"`
	Height              int  `json:"d"`
	ArgbBackgroundColor *int `json:"e"`
}

// ImageEventInfoDs @SerialName("c")
type ImageEventInfoDs struct {
	ImageEventInfo

	Dx1                 int  `json:"a"`
	Dy1                 int  `json:"b"`
	Dx2                 int  `json:"c"`
	Dy2                 int  `json:"d"`
	Sx1                 int  `json:"e"`
	Sy1                 int  `json:"f"`
	Sx2                 int  `json:"g"`
	Sy2                 int  `json:"h"`
	ArgbBackgroundColor *int `json:"i"`
}

// ImageEventInfoTransformed @SerialName("d")
type ImageEventInfoTransformed struct {
	ImageEventInfo

	Tx []float64 `json:"a"`
}

func ToImageEventInfo(i []interface{}) ImageEventInfo {
	t := i[0].(string)
	c := i[1].(map[string]interface{})

	switch t {
	case "a":
		return &ImageEventInfoXy{
			X:                   Jint(c["a"]),
			Y:                   Jint(c["b"]),
			ArgbBackgroundColor: Jnint(c["c"]),
		}
	case "b":
		return &ImageEventInfoXyWh{
			X:                   Jint(c["a"]),
			Y:                   Jint(c["b"]),
			Width:               Jint(c["c"]),
			Height:              Jint(c["d"]),
			ArgbBackgroundColor: Jnint(c["e"]),
		}
	case "c":
		return &ImageEventInfoDs{
			Dx1:                 Jint(c["a"]),
			Dy1:                 Jint(c["b"]),
			Dx2:                 Jint(c["c"]),
			Dy2:                 Jint(c["d"]),
			Sx1:                 Jint(c["e"]),
			Sy1:                 Jint(c["f"]),
			Sx2:                 Jint(c["g"]),
			Sy2:                 Jint(c["h"]),
			ArgbBackgroundColor: Jnint(c["i"]),
		}
	case "d":
		return &ImageEventInfoTransformed{
			Tx: Jaf64(c["a"]),
		}
	default:
		panic("invalid image id type")
	}
}
