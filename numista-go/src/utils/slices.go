package utils

func Contains(choices []string, value string) bool {
	for _, choice := range choices {
		if choice == value {
			return true
		}
	}
	return false
}
