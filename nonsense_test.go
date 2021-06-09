package nonsense

import "testing"

var SliceContainsStringTests = []struct {
	sl       []string
	str      string
	expected bool
}{
	{[]string{"abc", "avv", "sdgefgdf"}, "abc", true},
	{[]string{"abc", "avv", "sdgefgdf"}, "bcd", false},
	{[]string{"abc", "avv", "sdgefgdf"}, "", false},
	{[]string{}, "abc", false},
}

func TestSliceContainsString(t *testing.T) {
	for _, ct := range SliceContainsStringTests {
		if SliceContainsString(ct.sl, ct.str) != ct.expected {
			t.Errorf("SliceContainsString(%v, %s) = %v, want %v", ct.sl, ct.str, !ct.expected, ct.expected)
		}
	}
}
