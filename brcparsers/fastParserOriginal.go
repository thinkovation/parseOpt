package brcparsers

// FastParserOriginal is the version created at https://github.com/shraddhaag/1brc/blob/main/main.go
func FastParserOriginal(input string) (int64, error) {
	var isNegativeNumber bool
	if input[0] == '-' {
		isNegativeNumber = true
		input = input[1:]
	}
	var output int64
	switch len(input) {
	case 3:
		output = int64(input[0])*10 + int64(input[2]) - int64('0')*11
	case 4:
		output = int64(input[0])*100 + int64(input[1])*10 + int64(input[3]) - (int64('0') * 111)
	}

	if isNegativeNumber {
		return -output, nil
	}
	return output, nil
}
