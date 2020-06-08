package ascii

import "testing"

func TestStandardConverter(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{0, "`"},
		{1<<16 - 1, "$"},
	}

	for _, c := range cases {
		got := StandardConverter.ToAscii(c.in)
		if got != c.want {
			t.Errorf("StandardConverter.ToAscii(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
