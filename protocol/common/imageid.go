package common

import . "Projector/protocol/util"

type ImageId interface {
	imageId()
}

// ImageIdBuffered `json:"a"`
type ImageIdBuffered struct {
	ImageId

	RasterDataBufferSize int `json:"a"`
	ContentHash          int `json:"b"`
}

// ImageIdPVolatile `json:"b"`
type ImageIdPVolatile struct {
	ImageId

	Id int64 `json:"a"`
}

// ImageIdUnknown `json:"c"`
type ImageIdUnknown struct {
	ImageId

	ClassName string `json:"a"`
}

func ToImageId(i []interface{}) ImageId {
	t := i[0].(string)
	c := i[1].(map[string]interface{})

	switch t {
	case "a":
		return &ImageIdBuffered{
			RasterDataBufferSize: Jint(c["a"]),
			ContentHash:          Jint(c["b"]),
		}
	case "b":
		return &ImageIdPVolatile{
			Id: Jint64(c["a"]),
		}
	case "c":
		return &ImageIdUnknown{
			ClassName: c["a"].(string),
		}
	default:
		panic("invalid image id type")
	}
}
