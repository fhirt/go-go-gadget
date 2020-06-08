package ascii

import (
	"image/color"
)

var AverageConverter *averageConverter = new(averageConverter)

// averageConverter converts color values to brightness by averaging all values
type averageConverter struct {
}

// ToBrightness transforms color values by averaging to a matrix of brightness values
func (p *averageConverter) ToBrightness(color color.Color) int {
	r, g, b, _ := color.RGBA()
	avg := int((r + g + b) / 3)
	return avg
}

var MinMaxConverter *minMaxConverter = new(minMaxConverter)

type minMaxConverter struct {
}

func (p *minMaxConverter) ToBrightness(color color.Color) int {
	return (max(color) + min(color)) / 2
}

func max(color color.Color) int {
	r, g, b, _ := color.RGBA()
	max := b
	if r > g {
		if r > b {
			max = r
		}
	} else {
		if g > b {
			max = g
		}
	}
	return int(max)
}

func min(color color.Color) int {
	r, g, b, _ := color.RGBA()
	min := b
	if r < g {
		if r < b {
			min = r
		}
	} else {
		if g < b {
			min = g
		}
	}
	return int(min)
}

var LuminosityConverter *luminosityConverter = new(luminosityConverter)

type luminosityConverter struct {
}

func (p *luminosityConverter) ToBrightness(color color.Color) int {
	r, g, b, _ := color.RGBA()
	R, G, B := float32(r), float32(g), float32(b)

	return int(0.21*R + 0.72*G + 0.07*B)
}
