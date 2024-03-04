package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"parse/brcparsers"
	"strconv"
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
type FunctionStatsCollection struct {
	Name     string
	StatsCol []Stats
}
type StatsCollection struct {
	Functions []FunctionStatsCollection
}

func (s *StatsCollection) AddStats(name string, stats Stats) {
	for i, f := range s.Functions {
		if f.Name == name {
			s.Functions[i].StatsCol = append(s.Functions[i].StatsCol, stats)
			return
		}
	}
	var sc []Stats
	sc = append(sc, stats)
	s.Functions = append(s.Functions, FunctionStatsCollection{Name: name, StatsCol: sc})
}

func (s StatsCollection) WritetoCSV(filename string) error {
	// Create the CSV file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row
	header := []string{"FunctionName", "Mean", "Mode", "Min", "Max"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write rows for each function
	for _, funcStats := range s.Functions {
		for _, stats := range funcStats.StatsCol {
			row := []string{
				funcStats.Name,
				strconv.FormatFloat(stats.Mean, 'f', -1, 64),
				strconv.FormatInt(stats.Mode, 10),
				strconv.FormatInt(stats.Min, 10),
				strconv.FormatInt(stats.Max, 10),
			}
			if err := writer.Write(row); err != nil {
				return err
			}
		}
	}

	return nil
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
	//generateRandomNumbersToFile("random_numbers5000000.txt", 5000000)

	//	stringArray, err := readLinesFromFile("random_numbers1000.txt")
	//	stringArray, err := readLinesFromFile("random_numbers1000000.txt")
	//stringArray, err := readLinesFromFile("random_numbers100000000.txt")
	stringArray, err := readLinesFromFile("random_numbers5000000.txt")

	if err != nil {
		log.Fatal(err)
	}
	var functionsToCheck []FunctionInfo

	functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "FastParserOriginal", Function: brcparsers.FastParserOriginal})
	functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "FastParserBareBones", Function: brcparsers.FastParserBareBones})
	functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "FastParserBareBonesSwitch", Function: brcparsers.FastParserBareBonesSwitch})
	functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "FastParserOptimised", Function: brcparsers.FastParserOptimised})
	functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "CustomBRCParserChatGPT", Function: brcparsers.CustomBRCParserChatGPT})
	functionsToCheck = append(functionsToCheck, FunctionInfo{Name: "StandardParseFloat", Function: brcparsers.StandardParseFloat})

	//FastParserBareBones
	var overallSats StatsCollection
	loopcount := 100
	outerloopcount := 100
	for i := 0; i < outerloopcount; i++ {
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
			overallSats.AddStats(v.Name, v.Statistics)

		}

	}
	overallSats.WritetoCSV("stats2.csv")

}
