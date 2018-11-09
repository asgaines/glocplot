package conversions

import "github.com/asgaines/glocplot/errs"

func Map(point, fromMin, fromMax, toMin, toMax float64) (float64, error) {
	if toMin == toMax {
		return 0, errs.NewDivideByZeroError("Values for map destination must be different from one another")
	}

	perc := (point - fromMin) / (fromMax - fromMin)
	return perc*(toMax-toMin) + toMin, nil
}
