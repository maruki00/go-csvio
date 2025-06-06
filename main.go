package main

import (
	"fmt"
)

func main() {

	csv, err := NewReader("./file.csv", []byte{','})

	if err != nil {
		panic(err)
	}

	// fmt.Println(csv._headers)
	line, err := csv.lines()
	if err != nil {
		panic(err)
	}

	for l := range line {
		fmt.Println(l)
	}
}
