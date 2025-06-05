package gocsvio

import "os"

type CSV struct {
	_headers map[string]int
	line     []string
	fd       *os.File
}

func NewReader()
func (_this *CSV) Get(key string) any {
	index, ok := _this._headers[key]
	if !ok {
		return ""
	}
	return _this.line[index]
}
