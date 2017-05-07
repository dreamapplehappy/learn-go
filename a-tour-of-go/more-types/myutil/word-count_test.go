package myutil

import "testing"

func TestWordCount(t *testing.T) {
	cases := []struct {
		in   string;
		want map[string]int
	}{
		{"hello world", map[string]int{"hello world": 2}},
		{"hello world hi", map[string]int{"hello world hi": 3}},
	}

	for _, c := range cases {
		got := WordCount(c.in)
		if got[c.in] != c.want[c.in] {
			t.Error("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

