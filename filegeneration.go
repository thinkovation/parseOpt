package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// GenFile generates random numbers and writes them to a file.
func GenFile() {
	// Generate 10 random numbers and write to file
	err := generateRandomNumbersToFile("random_numbers100000000.txt", 100000000)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Random numbers generated and saved to file.")
}

// generateRandomNumbersToFile generates random numbers and writes them to a file.
func generateRandomNumbersToFile(filename string, count int) error {
	// Open file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Generate random numbers and write to file
	for i := 0; i < count; i++ {
		// Generate random number between -999 and 999
		randomNumber := rand.Intn(1999) - 999
		// Convert to float64 and divide by 10 to get one decimal place
		floatNumber := float64(randomNumber) / 10.0
		// Convert to string with one decimal place
		numberString := strconv.FormatFloat(floatNumber, 'f', 1, 64)
		// Write to file
		_, err := fmt.Fprintln(file, numberString)
		if err != nil {
			return err
		}
	}

	return nil
}

// readLinesFromFile reads lines from a file and returns them as a slice of strings.
func readLinesFromFile(filename string) ([]string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file
	for scanner.Scan() {
		// Read the line
		line := scanner.Text()

		// Add the line to the slice
		lines = append(lines, line)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
