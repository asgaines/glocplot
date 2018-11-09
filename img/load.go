package img

import (
	"image"
	"image/draw"
	"os"
)

func Load(imgFilename string) (*draw.Image, error) {
	imgFile, err := os.Open(imgFilename)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	var dimg draw.Image
	dimg, ok := img.(draw.Image)
	if !ok {
		// Try converting image to RGBA for drawing purposes
		newImg := image.NewRGBA(img.Bounds())
		draw.Draw(newImg, newImg.Bounds(), img, img.Bounds().Min, draw.Src)
		dimg = newImg
	}

	return &dimg, nil
}
