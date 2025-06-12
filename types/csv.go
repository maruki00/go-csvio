
package main


type CSV struct {
	_headers  map[string]int
	delimiter byte
	currPos   int64
	pathFile  string
}

