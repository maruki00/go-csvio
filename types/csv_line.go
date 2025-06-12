package types

type CSVLine struct {
	headers map[string]int
	Cols    []string
}

func NewCSVLine(headers map[string]int) *CSVLine {
	return &CSVLine{
		headers: headers,
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
