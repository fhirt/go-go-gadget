package ascii

import (
	"image/color"
	"testing"
)

func TestAverageConverter(t *testing.T) {
	cases := []struct {
		in   color.Color
		want int
	}{
		{color.RGBA{0, 0, 0, 0}, 0},
		{color.RGBA{100, 100, 100, 0}, 25700},
		{color.RGBA{255, 150, 0, 0}, 34695},
		{color.RGBA{80, 120, 0, 0}, 17133},
	}

	for _, c := range cases {
		got := AverageConverter.ToBrightness(c.in)
		if got != c.want {
			t.Errorf("AverageConverter.ToBrightness(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
