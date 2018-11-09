package plot

import (
	"image/color"
	"image/draw"
	"math"
)

func DrawLine(dimg *draw.Image, col color.RGBA, width uint, x0 float64, y0 float64, x1 float64, y1 float64) {
	dx := x1 - x0
	dy := y1 - y0

	// Protect against division by 0
	if dx == 0 {
		dx = 0.00000000001
	}

	dOff := math.Abs(dy / dx)
	off := 0.0

	y := y0

	for x := x0; x < x1; x++ {
		if dx >= dy {
			for i := -int(width / 2); i < int((width+1)/2); i++ {
				(*dimg).Set(int(x), int(y)+i, col)
			}
		} else {
			for i := -int(width / 2); i < int((width+1)/2); i++ {
				(*dimg).Set(int(x)+i, int(y), col)
			}
		}

		off = off + dOff
		if off >= 0.5 {
			if dy >= 0 {
				y++
			} else {
				y--
			}
			off = off - 1
		}
	}
}
