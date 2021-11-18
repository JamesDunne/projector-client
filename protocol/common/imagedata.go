package common

type ImageData interface {
	imageData()
}

// ImageDataPngBase64 @SerialName("a")
type ImageDataPngBase64 struct {
	ImageData

	PngBase64 string `json:"a"`
}

// ImageDataEmpty @SerialName("b")
type ImageDataEmpty struct {
	ImageData
}

func ToImageData(i []interface{}) ImageData {
	t := i[0].(string)
	c := i[1].(map[string]interface{})

	switch t {
	case "a":
		return &ImageDataPngBase64{
			PngBase64: c["a"].(string),
		}
	case "b":
		return &ImageDataEmpty{}
	default:
		panic("invalid image data type")
	}
}
