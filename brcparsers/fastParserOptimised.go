package brcparsers

// FastParserOptimised adapts the version created at https://github.com/shraddhaag/1brc/blob/main/main.go by swapping the
// switch out for an if else
func FastParserOptimised(input string) (int64, error) {
	// Variable to track if the number is negative
	var isNegativeNumber bool

	// Check if the input string starts with a negative sign
	if input[0] == '-' {
		isNegativeNumber = true
		// Remove the negative sign from the input
		input = input[1:]
	}

	// Variable to store the parsed integer value
	var output int64

	// Determine the length of the input string to decide the parsing logic
	if len(input) == 3 {
		// If the input string has length 3, parse accordingly
		// Parse the first and third characters and subtract the ASCII value of '0' multiplied by 11
		output = int64(input[0])*10 + int64(input[2]) - int64('0')*11
	} else {
		// If the input string has length other than 3, parse accordingly
		// Parse the first, second, and fourth characters and subtract the ASCII value of '0' multiplied by 111
		output = int64(input[0])*100 + int64(input[1])*10 + int64(input[3]) - (int64('0') * 111)
	}

	// If the number was negative, return the negation of the parsed value
	if isNegativeNumber {
		return -output, nil
	}

	// Otherwise, return the parsed value
	return output, nil
}
