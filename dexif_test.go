package dexif

import "testing"

func Benchmark_dexif(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dexif("./src.jpg", "dest.jpg")
	}
}
