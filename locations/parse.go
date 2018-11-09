package locations

import (
	"encoding/json"
	"os"

	"github.com/asgaines/glocplot/models"
)

func Parse(contents *os.File) (*models.Data, error) {
	var data models.Data

	decoder := json.NewDecoder(contents)
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
