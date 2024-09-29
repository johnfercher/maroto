package image

type Image struct {
	SourceKey string
}

func NewImage() *Image {
	return &Image{}
}
