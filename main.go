package main

import (
	"flag"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/asgaines/glocplot/img"
	"github.com/asgaines/glocplot/locations"
	"github.com/asgaines/glocplot/plot"
)

func main() {
	var locationF string
	var imageF string
	var outputImageF string
	var size uint
	var showLines bool

	flag.StringVar(&locationF, "location", "", "Path to file with location history")
	flag.StringVar(&imageF, "image", "./assets/images/Earth.png", "Path to map image")
	flag.StringVar(&outputImageF, "output", "result.png", "Image output filename")
	flag.UintVar(&size, "size", 1, "Size of each point and width of line")
	flag.BoolVar(&showLines, "lines", true, "Show lines between points for continuity")
	flag.Parse()

	if locationF == "" {
		flag.PrintDefaults()
		log.Fatal("Please provide path to location history file")
	}

	data, err := locations.Load(locationF)
	if err != nil {
		log.Fatal(err)
	}

	dimg, err := img.Load(imageF)
	if err != nil {
		log.Fatal(err)
	}

	err = plot.Locations(dimg, &data.Locations, size, showLines)
	if err != nil {
		log.Fatal(err)
	}

	err = img.Write(outputImageF, dimg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Successfully plotted %d locations to %s", len(data.Locations), outputImageF)
}
