package brcparsers

import "strconv"

func StandardParseFloat(input string) (int64, error) {

	f, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err // Return an error if conversion fails
	}

	return int64(f * 10), nil

}
