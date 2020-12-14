package dexif

import (
	"bytes"
	"encoding/binary"
	"errors"
)

var ErrExifMarkerNotFound = errors.New("can't find exif marker on file")

// JPEG is a struct that contains jpeg file bytes inside
type JPEG struct {
	buf []byte
}

// NewJPEG gives you a new JPEG, gets a parameter named buf
// don't use NewJPEG if you don't know exactly what buf contains.
// so use it when you exactly know the buf is a jpeg file.
func NewJPEG(buf []byte) *JPEG {
	return &JPEG{
		buf,
	}
}

func (jpeg *JPEG) Strip() ([]byte, error) {
	if jpeg.buf[2] == 0xFF {
		switch jpeg.buf[3] {
		case 0xE0:
			jfifSize := bytesToInt16(jpeg.buf[4:6])
			if jpeg.buf[jfifSize+4] == 0xFF && jpeg.buf[jfifSize+5] == 0xE1 {
				return stripExif(jpeg.buf, int(jfifSize+4))
			}
			break
		case 0xE1:
			return stripExif(jpeg.buf, 2)
		}
	}
	return nil, ErrExifMarkerNotFound
}

func stripExif(b []byte, exifMarkOffset int) ([]byte, error) {
	exifSize := bytesToInt16(b[exifMarkOffset+2 : exifMarkOffset+4])
	newImageSize := len(b) - int(exifSize)
	newBuf := make([]byte, 0, newImageSize)
	buffer := bytes.NewBuffer(newBuf)
	sum := exifMarkOffset + int(exifSize) + 2

	_, err := buffer.Write([]byte{0xFF, 0xD8})
	if err != nil {
		return nil, err
	}
	_, err = buffer.Write(b[sum:])
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func bytesToInt16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}
