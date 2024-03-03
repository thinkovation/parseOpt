package brcparsers

// FastParserOptimised adapts the version created at https://github.com/shraddhaag/1brc/blob/main/main.go by swapping the
// switch out for an if else
func FastParserBareBones(input string) (int64, error) {
	var isNegativeNumber bool
	if input[0] == '-' {
		isNegativeNumber = true
		input = input[1:]
	}
	var output int64
	if len(input) == 3 {
		//So the original calc was - where we essentially get to the
		//numerical value of the character by subtracting 48 (the byte value of zero) from it.
		//But this adds brackeds and a subtraction
		//So we can factor out the two subtractions of 48 by doing the multiple and subtracting
		//48*10 and 48*1 from the result.
		//48*11 = 528 so the line below... becomes the following line
		//output = int64((input[0]-48)*10) + int64(input[2]-48)
		output = int64(input[0]*10) + int64(input[2]) - 528
	} else {
		//Following the above - we've got three digits - so we subtract (48 * 100) + (48*10) + (48*1) from the number.
		//output = int64((input[0]-48)*100) + int64((input[1]-48)*10) + int64(input[3]-48)
		output = int64(input[0]*100) + int64(input[1]*10) + int64(input[2]) - 5328

	}

	if isNegativeNumber {
		return -output, nil
	}
	return output, nil
}
