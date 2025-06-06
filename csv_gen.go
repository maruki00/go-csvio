package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func _main() {
	const numRows = 10_000_000 // 10 million rows
	const filename = "large_file.csv"

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Printf("Generating %d rows into %s...\n", numRows, filename)

	// Write header
	_, err = file.WriteString("ID,Timestamp\n")
	if err != nil {
		fmt.Printf("Error writing header: %v\n", err)
		return
	}

	for i := 1; i <= numRows; i++ {
		// Using current timestamp for variety
		timestamp := time.Now().Format(time.RFC3339Nano)
		row := strconv.Itoa(i) + "," + timestamp + "\n"
		_, err := file.WriteString(row)
		if err != nil {
			fmt.Printf("Error writing row %d: %v\n", i, err)
			return
		}
	}

	fmt.Printf("Successfully generated %d rows into %s\n", numRows, filename)
}
