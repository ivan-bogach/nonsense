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

var StringIsNumericTests = []struct {
	str      string
	expected bool
}{
	{"test", false},
	{"", false},
	{"0", true},
	{"0.0", true},
	{"0.00", true},
	{"-1", true},
	{"-1000000000.0", true},
	{"-7777.0545454545", true},
	{"3254545454540545454545", true},
}

func TestStringIsNumeric(t *testing.T) {
	t.Parallel()

	for _, ct := range StringIsNumericTests {
		if StringIsNumeric(ct.str) != ct.expected {
			t.Errorf("StringIsNumeric(%s) = %v, want %v", ct.str, !ct.expected, ct.expected)
		}
	}
}
