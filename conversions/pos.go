package conversions

func E7ToStandard(e7Pos int64) float64 {
	return float64(e7Pos) / float64(1e7)
}
