package brcparsers

func CustomBRCParserChatGPT(input string) (int64, error) {
	var isNegativeNumber bool

	if input[0] == '-' {
		isNegativeNumber = true
		input = input[1:]
	}

	var output int64

	for _, char := range input {
		if char == '.' {
			continue // Ignore decimal points
		}
		output = output*10 + int64(char-'0')
	}

	if isNegativeNumber {
		return -output, nil
	}

	return output, nil
}
