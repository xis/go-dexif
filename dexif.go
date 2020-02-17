package dexif

import (
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
	buffer := make([]byte, 64000)
	_, err = f.ReadAt(buffer, 0)
	if err != nil {
		return err
	}
	if buffer[2] == 0xFF && buffer[3] == 0xE1 {
		for x := 1; x < 64000; x++ {
			if buffer[x-1] == 0xFF && buffer[x] == 0xD9 {
				fout, err := os.Create(destpath)
				if err != nil {
					return err
				}
				fout.Write([]byte{0xFF, 0xD8})
				f.Seek(int64(x), 0)
				_, err = io.Copy(fout, f)
				if err != nil {
					return err
				}
				defer fout.Close()
				return nil
			}
		}
	}
	return errors.New("no exif data found")
}
