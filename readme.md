# Experiments in optimising the parsing of numbers for the 1 Billion Rows Challenge

## Is there benefit in creating a custom parser for the numbers in the 1BRC?
All of the fastest GoLang solutions to the 1BRC follow a pretty similar pattern..

1) Read big chunks of the file at once
2) Make use of GoRoutines to make maximum use of the processor cores

One of the areas that immediately cropped up as a potential point of optimisation was in parsing the temperature value.

The test data is pretty predictable - The temperature will be between -99.9 degrees and 99.9 degrees (negative values will be preceded by the minus sign, positive values will have no sign.). It will always have one decimal place.

So the temperature string will always be between 3 and 5 characters long - eg 1.1, -23.5.

Aside from the minus symbol and the decimal point, all other characters will be in the set [0,1,2,3,4,5,6,7,8,9] and it's assumed that they'll be represented by single byte ascii codes.

So it clearly looks like there'll be scope for creating a very tuned parser.

This project currently contains 6 parsing functions, and I can't think of any other options to try - but hey, if you come up with something then let me know!

### FastParserOrginal

This- came from one of the entries to the 1BRC ( a very good one as it happens - check it out here https://github.com/shraddhaag/1brc/blob/main/main.go). It uses a pretty straightforward technique - it takes advantage of the fact that the asci character codes for digits are nicely ordered, with zero being the first (which is a byte with a value of 48) and  then go upwards through the digits (1 is 49, 2 is 50, 3 is 51... 9 is 57). It also dodges floating point drama by shifting the number to the left to remove the decimal - so 9.9 becomes 99 and 99.9 becomes 999 - Nice easy, compact, integers. 

It also exploits the fact that the format of the incoming number string is very stable - If you catch a minus sign and then snip it from the bytearray you basically have only two formats to consider...  nn.n and n.n


```
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

```
### FastParserOptimised
So I looked at this - and wondered what difference replacing the switch with an if/else block would do - While there's broad agreement that for many conditions switch is much better, what if there are only 2? This manifesed itself as the slightly wishful thinking "fastParserOptimised" - which identical but the switch is replaced with an if/else block. 

### FastPartserBarebones / FastPartserBarebonesSwitch
I then wondered if reducing the amount of maths the function does would make a difference. Since in the line
```
output = int64(input[0])*100 + int64(input[1])*10 + int64(input[3]) - (int64('0') * 111)
```
(int64('0') * 111) is effevtively a constant, does rewriting the line to ...
```
output = int64(input[0]*100) + int64(input[1]*10) + int64(input[2]) - 5328
```
Make thins quicker - Or does the compiler do this anyway?

So there are two versions - One using if/else and the other using switch.

### CustomParserChatGPT
This was just for kicks - and I suspect that I could have iterated with chatGPT to come up with a solution. But to be fair with ChatGPT, its first stab did work - although it's not super fast.

### StandardParser
This is just an implementation using stronv.ParseFloat - and it's there just to benchmark.

```
func StandardParseFloat(input string) (int64, error) {
	// Parse the input string as a floating-point number with 64-bit precision
	f, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err // Return an error if conversion fails
	}

	// Convert the parsed floating-point number to an integer by multiplying it by 10
	return int64(f * 10), nil
}
```
## The Results

Well... I'd say that there's a LOT of mileage in using a custom parser in this case. But, the differences between the optimised techniques don't seem big and so far (Running on a cluttered windows laptop) the results from each benchmarking run appear to vary a fair bit, so I suspect my conclusion may be that the optimised functions are all more or less the same. 

I may re-run the benchmarks on a standalone linux box which might give more predictable results than my laptop which is running a billion apps in the b/g.

The data below is from stats.xlsx which contains a pivot that summarises the results of running the test harness with 100 * 100 iterations on random_numbers5000000.txt which contains 5,000,000 rows.

The times are in milliseconds and are the time taken to process the string array of 5m numbers. The csv file that contains the raw data is stats2.csv, and it's generated by the app.





| Function                  | Mean (ms)  | Mode (ms) | Min (ms) | Max (ms) |
|---------------------------|------------|-----------|----------|----------|
| FastParserOriginal        | 47.28      | 44.51     | 32.68    | 73.16    |
| FastParserBareBones       | 48.07      | 46.19     | 33.92    | 73.62    |
| FastParserBareBonesSwitch | 47.69      | 44.69     | 32.9     | 74.42    |
| FastParserOptimised       | 48.10      | 45.91     | 33.84    | 76.15    |
| CustomBRCParserChatGPT    | 68.81      | 65.48     | 52.07    | 97.95    |
| StandardParseFloat        | 165.25     | 154.57    | 134.5    | 235 -     |
