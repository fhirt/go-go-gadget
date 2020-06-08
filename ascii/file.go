package ascii

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
	"github.com/pkg/errors"
)

const (
	maxDim = 100
)

// OpenImage opens an image file at the given path
func OpenImage(path string) (image.Image, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot open file at "+path)
	}
	defer reader.Close()

	img, err := jpeg.Decode(reader)
	if err != nil {
		return nil, errors.Wrap(err, "Decoding of image failed")
	}

	resized := resize.Thumbnail(maxDim, maxDim, img, resize.NearestNeighbor)

	return resized, nil
}
