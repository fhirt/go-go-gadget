// Package ascii implements simple jpeg to ascii transformations
package ascii

import (
	"image"
	"image/color"
)

// ColorConverter converts color to a value of brightness
type ColorConverter interface {
	ToBrightness(color color.Color) int
}

// BrightnessConverter transforms a given level of brightness to an appropriate ascii character
type BrightnessConverter interface {
	ToAscii(brightness int) string
}

func MakeAsciiArt(img image.Image, colorConverter ColorConverter, brightnessConverter BrightnessConverter) [][]string {
	min := img.Bounds().Min
	max := img.Bounds().Max

	artWork := make([][]string, max.Y)

	for y := min.Y; y < max.Y; y++ {
		artWork[y] = make([]string, max.X)
		for x := min.X; x < max.X; x++ {
			brightness := colorConverter.ToBrightness(img.At(x, y))
			artWork[y][x] = brightnessConverter.ToAscii(brightness)
		}
	}

	return artWork
}
