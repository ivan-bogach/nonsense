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
	t.Parallel()

	for _, ct := range SliceContainsStringTests {
		if SliceContainsString(ct.sl, ct.str) != ct.expected {
			t.Errorf("SliceContainsString(%v, %s) = %v, want %v", ct.sl, ct.str, !ct.expected, ct.expected)
		}
	}
}

var SendStringToTelegramTests = []struct {
	str      string
	expected int
}{
	{"test", 200},
	{"", 400},
}

func TestSendStringToTelegram(t *testing.T) {
	t.Parallel()

	for _, ct := range SendStringToTelegramTests {
		sstt := SendStringToTelegram(ct.str)
		if sstt != ct.expected {
			t.Errorf("SendStringToTelegram(%s) = %d, want %v", ct.str, sstt, ct.expected)
		}
	}
}
