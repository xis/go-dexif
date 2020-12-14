package dexif

import "errors"

type Image struct {
	buf []byte
}

func NewImage(buf []byte) *Image {
	return &Image{
		buf,
	}
}

func (image *Image) Strip() ([]byte, error) {
	if image.buf[0] == 0xFF && image.buf[1] == 0xD8 {
		jpeg := NewJPEG(image.buf)
		return jpeg.Strip()
	}
	return nil, errors.New("file format is not supported")
}
