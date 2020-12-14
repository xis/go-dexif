package dexif

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestStrip(t *testing.T) {
	b, err := ioutil.ReadFile("./test_images/img_2.jpg")
	if err != nil {
		t.Error(err)
	}
	jpeg := NewJPEG(b)
	newImage, err := jpeg.Strip()
	if err != nil {
		t.Error(err)
	}
	file, err := os.Create("img2_output.jpg")
	defer file.Close()
	if err != nil {
		t.Error(err)
	}
	file.Write(newImage)
}

func BenchmarkStrip(b *testing.B) {
	buf, err := ioutil.ReadFile("./test_images/img_1.jpg")
	if err != nil {
		b.Error(err)
	}
	jpeg := NewJPEG(buf)
	for n := 0; n < b.N; n++ {
		_, err = jpeg.Strip()
		if err != nil {
			b.Error(err)
		}
	}
}
