package nonsense

// SliceContainsString reports whether slaice contains string.
func SliceContainsString(sl []string, st string) bool {
	for _, s := range sl {
		if s == st {
			return true
		}
	}
	return false
}
