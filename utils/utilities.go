package utils

func ExtractEntryFromError(errorMessage string) (entry string) {
	firstIndex := 0
	secondIndex := 0

	for index, value := range errorMessage {
		if string(value) == "'" {
			if firstIndex == 0 {
				firstIndex = index
			} else if secondIndex == 0 {
				secondIndex = index
				break
			}
		}
	}
	entry = errorMessage[firstIndex + 1:secondIndex]
	return
}
