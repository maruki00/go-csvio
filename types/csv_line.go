package types

import "github.com/maruki00/go-csvio/utils"

type CSVLine struct {
	headers map[string]int
	delimeter byte
	Cols    []string
}

func NewCSVLine(headers map[string]int, del byte) *CSVLine {
	return &CSVLine{
		headers: headers,
		delimeter: del,
		Cols:    make([]string, len(headers)),
	}
}

func (_this *CSVLine) Get(key string) any {
	index, ok := _this.headers[key]
	if !ok {
		return ""
	}
	return _this.Cols[index]
}

func (_this *CSVLine) parse(row string) {
	utils.Split(row, string(_this.delimeter) , &_this.Cols)
}

