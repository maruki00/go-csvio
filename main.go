package main

import "fmt"

func main() {
	csv, err := NewReader("./file.csv", []byte{','})
	if err != nil {
		panic(err)
	}
	lines, err := csv.yield()
	if err != nil {
		panic(err)
	}

	for line := range lines {
		fmt.Println(" -- ", line.Get("id"), line.Get("city"))
	}
}
