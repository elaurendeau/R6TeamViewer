package utils


func Contains(slice []string, element string) bool {
	for _, current := range slice {
		if current == element {
			return true
		}
	}
	return false
}
