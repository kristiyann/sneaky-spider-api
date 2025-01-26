package util

func StringPtr(s string) *string {
	return &s
}

func String(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func StringSliceContains(s []string, value string) bool {
	for _, item := range s {
		if item == value {
			return true
		}
	}

	return false
}
