package notes

import "testing"

func TestNote(t *testing.T) {
	var cases = []struct {
		from, to *Pitch
		expect   Interval
	}{
		{&Pitch{A, 3, Natural}, &Pitch{A, 3, Natural}, 0},
		{&Pitch{A, 3, Natural}, &Pitch{B, 3, Natural}, 2},
		{&Pitch{A, 3, Natural}, &Pitch{G, 3, Natural}, -2},
		{&Pitch{E, 3, Natural}, &Pitch{F, 3, Natural}, 1},
		{&Pitch{F, 3, Natural}, &Pitch{E, 3, Natural}, -1},
		{&Pitch{B, 3, Natural}, &Pitch{C, 4, Natural}, 1},
		{&Pitch{C, 4, Natural}, &Pitch{B, 3, Natural}, -1},
		{&Pitch{C, 2, Natural}, &Pitch{C, 4, Natural}, 24},
		{&Pitch{C, 3, Natural}, &Pitch{C, 3, Sharp}, 1},
		{&Pitch{C, 3, Flat}, &Pitch{B, 2, Natural}, 0},
		{&Pitch{B, 3, DoubleSharp}, &Pitch{C, 4, DoubleFlat}, -3},
	}

	for _, c := range cases {
		if got := c.from.Interval(c.to); got != c.expect {
			t.Errorf("(%v, %v) expected %v, got %v\n", c.from, c.to, c.expect, got)
		}
	}
}
