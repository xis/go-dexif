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
	b := make([]byte, 2)
	f.Seek(2, 0)
	f.Read(b)
	res := bytes.Compare(b, []byte{0xFF, 0xE1})
	if res == 0 {
		for i := 0; i < 32768; i++ {
			f.Read(b)
			res := bytes.Compare(b, []byte{0xFF, 0xD9})
			if res == 0 {
				break
			}
		}
	} else {
		return errors.New("no exif data found")
	}
	fout, _ := os.Create(destpath)
	fout.Write([]byte{0xFF, 0xD8})
	io.Copy(fout, f)
	defer fout.Close()
	return nil
}
