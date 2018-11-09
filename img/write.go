package img

import (
	"image/draw"
	"image/png"
	"os"
)

func Write(outputFile string, dimg *draw.Image) error {
	imgOutF, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer imgOutF.Close()

	return png.Encode(imgOutF, *dimg)
}
