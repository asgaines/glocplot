package plot

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/asgaines/glocplot/conversions"
	"github.com/asgaines/glocplot/models"
)

type point struct {
	x float64
	y float64
}

func Locations(dimg *draw.Image, locations *[]models.Location, size uint, showLines bool) error {
	red := color.RGBA{255, 50, 50, 255}
	blue := color.RGBA{150, 150, 255, 255}

	previous := point{
		x: -1,
		y: -1,
	}

	for _, location := range *locations {
		lat := conversions.E7ToStandard(location.LatitudeE7)
		long := conversions.E7ToStandard(location.LongitudeE7)

		y, err := conversions.Map(lat, 90, -90, 0, float64((*dimg).Bounds().Max.Y))
		if err != nil {
			return err
		}

		x, err := conversions.Map(long, -180, 180, 0, float64((*dimg).Bounds().Max.X))
		if err != nil {
			return err
		}

		if showLines {
			if previous.x != -1 && previous.y != -1 {
				var x0, y0, x1, y1 float64

				if x > previous.x {
					x0, y0 = previous.x, previous.y
					x1, y1 = x, y
				} else {
					x0, y0 = x, y
					x1, y1 = previous.x, previous.y
				}

				DrawLine(dimg, blue, size, x0, y0, x1, y1)
			}
		}

		draw.Draw(*dimg, image.Rect(
			int(x)-int(size/2),
			int(y)-int(size/2),
			int(x)+int((size+1)/2),
			int(y)+int((size+1)/2),
		), &image.Uniform{red}, image.ZP, draw.Src)

		previous = point{
			x: x,
			y: y,
		}
	}

	return nil
}
