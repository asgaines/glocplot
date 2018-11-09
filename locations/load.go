package locations

import (
	"os"

	"github.com/asgaines/glocplot/models"
)

func Load(locFile string) (*models.Data, error) {
	contents, err := os.Open(locFile)
	if err != nil {
		return nil, err
	}
	defer contents.Close()

	data, err := Parse(contents)
	if err != nil {
		return nil, err
	}

	// Locations are ordered in descending time. Reverse for forward time chronology
	for i, j := 0, len(data.Locations)-1; i < j; i, j = i+1, j-1 {
		data.Locations[i], data.Locations[j] = data.Locations[j], data.Locations[i]
	}

	return data, nil
}
