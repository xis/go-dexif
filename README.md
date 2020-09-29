
<div align="center">
  <h1>go-dexif</h1>
  
  [![forthebadge](https://forthebadge.com/images/badges/check-it-out.svg)](https://forthebadge.com)

removes exif, not safe currently
</div>

# install
```bash
go get github.com/xis/go-dexif
```

# usage
```go
import "github.com/xis/go-dexif"

func main() {
	err := dexif.Dexif("./test.jpeg", "dest.jpg")
	if err != nil {
	    panic(err)
	}
}
```
