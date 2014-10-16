package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

// Image is
type Image struct {
	H, W int
	C    uint8
}

// ColorModel returns the Image's color model.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.W, i.H)
}

// At returns the color of the pixel at (x, y).
// ex 0 0 0 0
func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.C, i.C, 255, 255}
}

func main() {
	m := Image{100, 200, 250}
	pic.ShowImage(m)
}
