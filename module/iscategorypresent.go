package module

// Function to check if a category is present in the slice
func isCategoryPresent(category string, categories []string) bool {
	for _, c := range categories {
		if c == category {
			return true
		}
	}
	return false
}
