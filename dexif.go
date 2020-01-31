package dexif

import (
	"bytes"
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
	b1 := make([]byte, 2)
	f.Seek(2, 0)
	f.Read(b1)
	res := bytes.Compare(b1, []byte{0xFF, 0xE1})
	if res == 0 {
		for i := 0; i < 32768; i++ {
			_, err := f.Read(b1)
			if err != nil {
				return err
			}
			res := bytes.Compare(b1, []byte{0xFF, 0xD9})
			if res == 0 {
				break
			}
		}
	} else {
		return errors.New("no exif data found")
	}
	fout, err := os.Create(destpath)
	if err != nil {
		return err
	}
	fout.Write([]byte{0xFF, 0xD8})
	_, err = io.Copy(fout, f)
	if err != nil {
		return err
	}
	defer fout.Close()
	return nil
}
