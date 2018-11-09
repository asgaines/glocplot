package conversions

import (
	"reflect"
	"testing"

	"github.com/asgaines/glocplot/errs"
)

func TestMap(t *testing.T) {
	cases := []struct {
		point    float64
		fromMin  float64
		fromMax  float64
		toMin    float64
		toMax    float64
		expected float64
		err      error
	}{
		{
			point:    5,
			fromMin:  0,
			fromMax:  10,
			toMin:    0,
			toMax:    100,
			expected: 50,
			err:      nil,
		},
		{
			point:    2.5,
			fromMin:  0,
			fromMax:  10,
			toMin:    0,
			toMax:    100,
			expected: 25,
			err:      nil,
		},
		{
			point:    15,
			fromMin:  0,
			fromMax:  10,
			toMin:    0,
			toMax:    100,
			expected: 150,
			err:      nil,
		},
		{
			point:    0,
			fromMin:  -10,
			fromMax:  10,
			toMin:    0,
			toMax:    15,
			expected: 7.5,
			err:      nil,
		},
		{
			point:    5,
			fromMin:  -10,
			fromMax:  10,
			toMin:    -100,
			toMax:    100,
			expected: 50,
			err:      nil,
		},
		{
			point:    -15,
			fromMin:  -20,
			fromMax:  -10,
			toMin:    10,
			toMax:    20,
			expected: 15,
			err:      nil,
		},
		{
			point:    -15,
			fromMin:  -20,
			fromMax:  -10,
			toMin:    -100,
			toMax:    -50,
			expected: -75,
			err:      nil,
		},
		{
			point:    75,
			fromMin:  70,
			fromMax:  80,
			toMin:    -80,
			toMax:    -70,
			expected: -75,
			err:      nil,
		},
		{
			point:    -10,
			fromMin:  0,
			fromMax:  -20,
			toMin:    0,
			toMax:    100,
			expected: 50,
			err:      nil,
		},
		{
			point:    10,
			fromMin:  20,
			fromMax:  0,
			toMin:    -50,
			toMax:    -100,
			expected: -75,
			err:      nil,
		},
		{
			point:    500,
			fromMin:  0,
			fromMax:  1000,
			toMin:    0,
			toMax:    -100,
			expected: -50,
			err:      nil,
		},
		{
			point:    45,
			fromMin:  90,
			fromMax:  -90,
			toMin:    0,
			toMax:    400,
			expected: 100,
			err:      nil,
		},
		{
			point:    5,
			fromMin:  5,
			fromMax:  5,
			toMin:    5,
			toMax:    5,
			expected: 0,
			err:      errs.ErrZeroDivision{},
		},
	}

	for _, c := range cases {
		result, err := Map(c.point, c.fromMin, c.fromMax, c.toMin, c.toMax)
		if reflect.TypeOf(c.err) != reflect.TypeOf(err) {
			t.Errorf("Expected error type %T, received %T", c.err, err)
		}
		if err == nil && result != c.expected {
			t.Errorf("Expected %v, got %v", c.expected, result)
		}
	}
}
