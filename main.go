package main

import (
	"fmt"
)

func main() {

	csv, err := NewReader("./file.csv", []byte{','})

	if err != nil {
		panic(err)
	}

	fmt.Println(csv._headers)
	line, _ := csv.yield()

	for l := range line {
		fmt.Println(l)
	}
	fmt.Println(<-line)
	fmt.Println(<-line)
	fmt.Println(<-line)
}
