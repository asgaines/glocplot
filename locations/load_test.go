package locations

import (
	"reflect"
	"testing"

	"github.com/asgaines/glocplot/models"
)

func TestLoad(t *testing.T) {
	data, err := Load("./testdata/locations.json")
	if err != nil {
		t.Error(err)
	}

	expected := &models.Data{
		Locations: []models.Location{
			models.Location{
				TimestampMs:      "1540237938443",
				LatitudeE7:       399861229,
				LongitudeE7:      -1070338490,
				Accuracy:         15,
				Altitude:         1504,
				VerticalAccuracy: 2,
				Activities: []models.ActivityWrapper{
					{
						TimestampMs: "1540237559418",
						Activities: []models.Activity{
							{
								Type:       "STILL",
								Confidence: 100,
							},
						},
					},
				},
			},
			models.Location{
				TimestampMs:      "1540237966433",
				LatitudeE7:       399861050,
				LongitudeE7:      -1070338568,
				Accuracy:         15,
				Altitude:         1504,
				VerticalAccuracy: 2,
				Activities: []models.ActivityWrapper{
					{
						TimestampMs: "1540438013894",
						Activities: []models.Activity{
							{
								Type:       "STILL",
								Confidence: 100,
							},
						},
					},
				},
			},
			models.Location{
				TimestampMs:      "1540238172660",
				LatitudeE7:       399861207,
				LongitudeE7:      -1070338646,
				Accuracy:         15,
				Altitude:         1504,
				VerticalAccuracy: 2,
				Activities: []models.ActivityWrapper{
					{
						TimestampMs: "1540438174470",
						Activities: []models.Activity{
							{
								Type:       "STILL",
								Confidence: 100,
							},
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, got %v", expected, data)
	}
}
