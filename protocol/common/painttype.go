package common

type PaintType int

const (
	PaintTypeDraw PaintType = iota
	PaintTypeFill
)

func ToPaintType(c string) PaintType {
	switch c {
	case "a":
		return PaintTypeDraw
	case "b":
		return PaintTypeFill
	default:
		panic("invalid paint type")
	}
}
