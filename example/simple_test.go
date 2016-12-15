package example

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		x, y, want int
	}{
		{4, 5, 9},
		{2, 3, 5},
		{3, 4, 7},
	}

	for _, c := range cases {
		got := Add(c.x, c.y)
		if got != c.want {
			t.Errorf("add(%d,%d) = %d, want %d", c.x, c.y, got, c.want)
		}
	}

}
