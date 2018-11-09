package conversions

import (
	"fmt"
	"time"
)

func TimestampToDate(timestampMs int64) string {
	t := time.Unix(timestampMs/1000, 0)
	year, month, day := t.Date()

	return fmt.Sprintf("%d-%d-%d", year, month, day)
}
