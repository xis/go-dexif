
<div align="center">
  <h1>go-dexif</h1>
  
  [![forthebadge](https://forthebadge.com/images/badges/check-it-out.svg)](https://forthebadge.com)

removes exif, not safe currently, have problems with images that contains orientation data, supports only jpeg.
</div>

# install
```bash
go get github.com/xis/go-dexif
```

# usage
```go
import (
	"log"
	"github.com/xis/go-dexif"
)

func main() {
	buf, err := ioutil.ReadFile("./image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	image := NewImage(buf)
	newImage, err := image.Strip()
	if err != nil {
		log.Fatal(err)
	}
}
```
