package common

type PathSegment interface {
	pathSegment()
}

// PathSegmentMoveTo `json:"a"`
type PathSegmentMoveTo struct {
	PathSegment

	Point Point `json:"a"`
}

// PathSegmentLineTo `json:"b"`
type PathSegmentLineTo struct {
	PathSegment

	Point Point `json:"a"`
}

// PathSegmentQuadTo `json:"c"`
type PathSegmentQuadTo struct {
	PathSegment

	Point1 Point `json:"a"`
	Point2 Point `json:"b"`
}

// PathSegmentCubicTo `json:"d"`
type PathSegmentCubicTo struct {
	PathSegment

	Point1 Point `json:"a"`
	Point2 Point `json:"b"`
	Point3 Point `json:"c"`
}

// PathSegmentClose `json:"e"`
type PathSegmentClose struct {
	PathSegment
}

type WindingType int

const (
	// WindingEvenOdd `json:"a"`
	WindingEvenOdd WindingType = iota
	// WindingNonZero `json:"b"`
	WindingNonZero
)

type CommonPath struct {
	Segments []PathSegment `json:"a"`
	Winding  WindingType   `json:"b"`
}
