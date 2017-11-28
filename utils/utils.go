package utils

func OnlyCyrylik(s string) bool {
	for _, c := range s {
		if c < 'А' || c > 'я' {
			return false
		}
	}
	return true
}
