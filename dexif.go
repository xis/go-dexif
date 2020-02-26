package dexif

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

// Dexif @ removes exif data in image
func Dexif(filepath string, destpath string) error {
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		return err
	}
	buffer := make([]byte, 4)
	_, err = f.ReadAt(buffer, 2)
	if err != nil {
		return err
	}

	if buffer[0] == 0xFF && buffer[1] == 0xE1 {
		exifSize := binary.BigEndian.Uint16(buffer[2:4])
		exifSize += 4
		fout, err := os.Create(destpath)
		defer fout.Close()
		if err != nil {
			return err
		}
		fout.Write([]byte{0xFF, 0xD8})
		f.Seek(int64(exifSize), 0)
		_, err = io.Copy(fout, f)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("exif data not found")
}
