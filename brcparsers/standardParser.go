package brcparsers

import "strconv"

// StandardParseFloat parses a string representation of a floating-point number into an integer value.
// It multiplies the parsed floating-point number by 10 and returns the resulting integer value.
func StandardParseFloat(input string) (int64, error) {
	// Parse the input string as a floating-point number with 64-bit precision
	f, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err // Return an error if conversion fails
	}

	// Convert the parsed floating-point number to an integer by multiplying it by 10
	return int64(f * 10), nil
}
