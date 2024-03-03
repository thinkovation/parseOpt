package main

import (
	"fmt"
	"log"
	"math"
	"parse/brcparsers"
	"time"
)

type FunctionInfo struct {
	Name       string
	Function   parseFunc
	Statistics Stats
}

type parseFunc func(string) (int64, error)

func processStringArrays(strings []string, fn parseFunc) (int64, error) {
	//fmt.Println("Processing strings...", len(strings))
	ts := time.Now()
	for _, s := range strings {
		_, err := fn(s)
		if err != nil {
			return 0, err
		}
	}
	//fmt.Println("Done")
	//elapsed := time.Since(ts).Nanoseconds()
	elapsed := time.Since(ts).Milliseconds()

	//fmt.Println(time.Since(ts).Microseconds())

	return elapsed, nil
}

type Stats struct {
	Mean float64
	Mode int64
	Min  int64
	Max  int64
}

func (s Stats) String() string {
	return fmt.Sprintf("Mean: %.2f, Mode: %d, Min: %d, Max: %d", s.Mean, s.Mode, s.Min, s.Max)
}
func (f FunctionInfo) String() string {
	return fmt.Sprintf("%s,%.2f,%d,%d,%d", f.Name, f.Statistics.Mean, f.Statistics.Mode, f.Statistics.Min, f.Statistics.Max)
}

// calculateStatistics calculates the mean, mode, minimum, and maximum values
// from an array of int64 and returns them.
func calculateStatistics(arr []int64) Stats {
	// Calculate mean
	sumt := int64(0)
	max := int64(0)
	min := int64(math.MaxInt64)
	mode := int64(0)

	freqMap := make(map[int64]int)
	maxFreq := 0

	for _, num := range arr {
		sumt += num
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
		freqMap[num]++
		if freqMap[num] > maxFreq {
			maxFreq = freqMap[num]
			mode = num
		}

	}
	mean := float64(sumt) / float64(len(arr))
	//fmt.Println("Sumt", sumt)
	//fmt.Println(len(arr))

	return Stats{
		Mean: mean,
		Mode: mode,
		Min:  min,
		Max:  max,
	}

}
func main() {
	// random_numbers1000000.txt
	//	generateRandomNumbersToFile("random_numbers1000.txt", 1000)

	//	stringArray, err := readLinesFromFile("random_numbers1000.txt")
	stringArray, err := readLinesFromFile("random_numbers1000000.txt")
	//stringArray, err := readLinesFromFile("random_numbers100000000.txt")

	if err != nil {
		log.Fatal(err)
	}
	var functionsToCheck []FunctionInfo

	functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "FastParserOriginal", Function: brcparsers.FastParserOriginal})
	functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "FastParserBareBones", Function: brcparsers.FastParserBareBones})
	functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "FastParserBareBonesSwitch", Function: brcparsers.FastParserBareBonesSwitch})

	//functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "FastParserOriginal", Function: brcparsers.FastParserOriginal})
	//functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "FastParserBareBones", Function: brcparsers.FastParserBareBones})
	//FastParserBareBonesSwitch
	//functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "FastParserOptimised", Function: brcparsers.FastParserOptimised})
	//functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "CustomBRCParserChatGPT", Function: brcparsers.CustomBRCParserChatGPT})
	//functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "StandardParseFloat", Function: brcparsers.StandardParseFloat})

	//FastParserBareBones
	loopcount := 300

	for k, v := range functionsToCheck {
		var times []int64
		fmt.Println("Function", v.Name)
		for i := 0; i < loopcount; i++ {

			elapsed, _ := processStringArrays(stringArray, v.Function)

			times = append(times, elapsed)

		}
		fmt.Println(times)
		v.Statistics = calculateStatistics(times)
		functionsToCheck[k].Statistics = v.Statistics

	}
	for _, v := range functionsToCheck {
		fmt.Println(v.String())
	}

}
