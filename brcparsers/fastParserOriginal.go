package brcparsers

// FastParserOriginal is a function that converts a string to an integer using ASCII arithmetic.
// It expects the input string to represent an integer number. It handles negative numbers and numbers with either 3 or 4 digits.
// If the input string represents a negative number, it returns the negative of the calculated value.
// It returns the converted integer value and nil if conversion is successful, or 0 and an error if conversion fails.
func FastParserOriginal(input string) (int64, error) {
	// Variable to track if the number is negative
	var isNegativeNumber bool

	// Check if the input string starts with a negative sign '-'
	if input[0] == '-' {
		isNegativeNumber = true
		// Remove the negative sign from the input string
		input = input[1:]
	}

	// Variable to store the converted integer value
	var output int64

	// Switch statement based on the length of the input string
	switch len(input) {
	case 3:
		// For 3-digit numbers, calculate the integer value using ASCII arithmetic
		output = int64(input[0])*10 + int64(input[2]) - int64('0')*11
	case 4:
		// For 4-digit numbers, calculate the integer value using ASCII arithmetic
		output = int64(input[0])*100 + int64(input[1])*10 + int64(input[3]) - (int64('0') * 111)
	}

	// If the number was negative, return the negative of the calculated value
	if isNegativeNumber {
		return -output, nil
	}

	// Return the converted integer value and nil, indicating success
	return output, nil
}
