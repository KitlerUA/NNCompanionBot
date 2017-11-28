package utils

func OnlyCyrylik(s string) bool {
	for _, c := range s {
		if c < 'а' || c > 'Я' {
			return false
		}
	}
	return true
}
