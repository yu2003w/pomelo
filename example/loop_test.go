package example

import (
	"testing"
)

func TestCalpow(t *testing.T) {
	cases := []struct {
		x, n, lim, want float64
	}{
		{3, 2, 10, 9},
		{3, 3, 20, 20},
	}

	for _, c := range cases {
		got := Calpow(c.x, c.n, c.lim)
		if got != c.want {
			t.Errorf("Calpow(%g,%g,%g) = %g, want %g", c.x, c.n, c.lim, got, c.want)
		}
	}
}

func TestForsim(t *testing.T) {
	cases := []struct {
		num, want int
	}{
		{-1, -1},
		{-2, -2},
		{2, 3},
	}

	for _, c := range cases {
		got := Forsim(c.num)
		if got != c.want {
			t.Errorf("Forsim(%d) = %d, want %d", c.num, got, c.want)
		}
	}
}

func TestForloop(t *testing.T) {
	cases := []struct {
		num, want int
	}{
		{10, 45},
		{11, 55},
	}

	for _, c := range cases {
		got := Forloop(c.num)
		if got != c.want {
			t.Errorf("Forloop(%d) = %d, want %d", c.num, got, c.want)
		}
	}

}
