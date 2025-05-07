package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	// FIXME: hardcoded values these could be passes or read from the image struct
	// function we're using to produce the image
	v := uint8(x * y)
	return color.RGBA{v, v, 255, 255}
}

func (i Image) Bounds() image.Rectangle {
	// FIXME: hardcoded values these could be passes or read from the image struct
	return image.Rect(0, 0, 100, 100)
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
