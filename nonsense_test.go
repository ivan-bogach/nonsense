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
		sstt, _ := SendStringToTelegram(ct.str)
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

var CompareStringSlicesTests = []struct {
	sl1      []string
	sl2      []string
	expected bool
}{
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"abc", "avv", "sdgefgdf"}, true},
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"avv", "abc", "sdgefgdf"}, false},
	{[]string{"abc", "avv", "sdgefgdf"}, []string{}, false},
	{[]string{}, []string{"abc", "avv", "sdgefgdf"}, false},
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"abce", "avv", "sdgefgdf"}, false},
	{nil, []string{"abce", "avv", "sdgefgdf"}, false},
	{nil, nil, true},
}

func TestCompareStringSlices(t *testing.T) {
	t.Parallel()

	for _, ct := range CompareStringSlicesTests {
		if CompareStringSlices(ct.sl1, ct.sl2) != ct.expected {
			t.Errorf("CompareStringSlices(%v, %v) = %v, want %v", ct.sl1, ct.sl2, !ct.expected, ct.expected)
		}
	}
}

var StringSlicesIntersectionTests = []struct {
	sl1      []string
	sl2      []string
	expected []string
}{
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"abc", "avv", "sdgefgdf"}, []string{"abc", "avv", "sdgefgdf"}},
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"avv", "abc", "sdgefgdf"}, []string{"avv", "abc", "sdgefgdf"}},
	{[]string{"abc", "avv", "sdgefgdf"}, []string{}, nil},
	{[]string{}, []string{"abc", "avv", "sdgefgdf"}, nil},
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"avv", "abce", "sdgefgdf"}, []string{"avv", "sdgefgdf"}},
	{[]string{"sdfsdfsdfsd", "abc", "sdfgrtryjdgh", "avv", "asdfgsdf", "sdgefgdf"}, []string{"avv", "edgdf", "abce", "asdsdfgsdf", "sdgefgdf"}, []string{"avv", "sdgefgdf"}},
	{nil, []string{"abce", "avv", "sdgefgdf"}, nil},
	{nil, nil, nil},
}

func TestStringSlicesIntersection(t *testing.T) {
	t.Parallel()

	for _, ct := range StringSlicesIntersectionTests {
		ssi := StringSlicesIntersection(ct.sl1, ct.sl2)
		if !CompareStringSlices(ssi, ct.expected) {
			t.Errorf("StringSlicesIntersection(%v, %v) = %v, want %v", ct.sl1, ct.sl2, ssi, ct.expected)
		}
	}
}

var OnlyUniqueTests = []struct {
	sl1      []string
	expected []string
}{
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"abc", "avv", "sdgefgdf"}},
	{[]string{"sdgefgdf", "abc", "avv", "sdgefgdf"}, []string{"abc", "avv", "sdgefgdf"}},
	{[]string{}, []string{}},
}

func TestOnlyUnique(t *testing.T) {
	t.Parallel()

	for _, ct := range OnlyUniqueTests {
		ou := OnlyUnique(ct.sl1)
		if !CompareStringSlices(ou, ct.expected) {
			t.Errorf("OnlyUnique(%v) = %v, want %v", ct.sl1, ou, ct.expected)
		}
	}
}

var StringSlicesUnionTests = []struct {
	sl1      []string
	sl2      []string
	expected []string
}{
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"abasfc", "avsdfsdfv", "sdgefgdf"}, []string{"abasfc", "abc", "avsdfsdfv", "avv", "sdgefgdf"}},
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"avv", "abc", "sdgefgdf"}, []string{"abc", "avv", "sdgefgdf"}},
	{[]string{"abc", "avv", "sdgefgdf"}, []string{}, []string{"abc", "avv", "sdgefgdf"}},
	{[]string{}, []string{"abc", "avv", "sdgefgdf"}, []string{"abc", "avv", "sdgefgdf"}},
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"avv", "abce", "sdgefgdf"}, []string{"abc", "abce", "avv", "sdgefgdf"}},
	{nil, []string{"abce", "avv", "sdgefgdf"}, []string{"abce", "avv", "sdgefgdf"}},
}

func TestStringSlicesUnion(t *testing.T) {
	t.Parallel()

	for _, ct := range StringSlicesUnionTests {
		ssu := StringSlicesUnion(ct.sl1, ct.sl2)
		if !CompareStringSlices(ssu, ct.expected) {
			t.Errorf("StringSlicesUnion(%v, %v) = %v, want %v", ct.sl1, ct.sl2, ssu, ct.expected)
		}
	}
}

var StringSliceDifferenceTests = []struct {
	sl1      []string
	sl2      []string
	expected []string
}{
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"abasfc", "avsdfsdfv", "sdgefgdf"}, []string{"abc", "avv"}},
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"avv", "abc", "sdgefgdf"}, nil},
	{[]string{"abc", "avv", "sdgefgdf"}, []string{}, []string{"abc", "avv", "sdgefgdf"}},
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"sd", "sdsdsdfsdfdfsdf"}, []string{"abc", "avv", "sdgefgdf"}},
	{[]string{}, []string{"abc", "avv", "sdgefgdf"}, nil},
	{[]string{"abc", "avv", "sdgefgdf"}, []string{"avv", "abce", "sdgefgdf"}, []string{"abc"}},
	{nil, []string{"abce", "avv", "sdgefgdf"}, nil},
}

func TestStringSliceDifference(t *testing.T) {
	t.Parallel()

	for _, ct := range StringSliceDifferenceTests {
		ssd := StringSliceDifference(ct.sl1, ct.sl2)
		if !CompareStringSlices(ssd, ct.expected) {
			t.Errorf("StringSliceDifference(%v, %v) = %v, want %v", ct.sl1, ct.sl2, ssd, ct.expected)
		}
	}
}
