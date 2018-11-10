# glocplot: Google Location Plotter

Plot your Google location data on a map for visualization. All processing performed on the executing machine; your data is entirely private to you (and Google, of course...).

![Example result](assets/images/Example.png)
Example result (Download image to view more detail)

![Desktop wallpaper](assets/images/Wallpaper.png)
Great for desktop wallpaper, especially after cropping.

## Installation

`go get github.com/asgaines/glocplot`\
`go install github.com/asgaines/glocplot`

## Usage

First download data from Google ([google.com/takeout](https://google.com/takeout)). Ensure the `Location History` app is toggled on and `JSON format` is selected. Download and unzip.

Example ``glocplot -location ~/Downloads/Takeout/Location\ History/Location\ History.json -image `go env GOPATH`/src/github.com/asgaines/glocplot/assets/images/Earth.png``
Output will be written to `result.png` in working directory.

Arguments
- `location` Path to file with location history data downloaded from Google Takeout (`Location History.json`)
- `image` Path to image on which your location data is plotted. A high-res, centered image is provided at `assets/images/Earth.png`
- `output` Output location. Defaults to `result.png`
- `size` Width of points and lines plotted to image. The smaller the width, the more the detail.
- `lines` Whether to display lines connecting location points.

Open up the image and enjoy!
