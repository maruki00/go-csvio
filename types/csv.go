package types

type CSV struct {
	_headers  map[string]int
	delimiter byte
	currPos   int64
	pathFile  string
}

func NewCSV(csvpath string) (*CSV, error) {
	_this := &CSV{
		_headers:   make(map[string]int),
		delimiter: '',
		csvPath:   csvpath,
	}
	if err := _this.parseHeader(); err != nil {
		return nil, err
	}
	return _this, nil
}

func (_this *CSV)extractDelimeter(){

}

func (_this *CSV) SetDelimiters(delimiters []byte) {
	_this.delimiters = delimiters
}

func (_this *CSV) SeekToLine(file *os.File, line uint) error {
	var buf [1]byte
	var offset int64 = 0
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	for line > 0 {
		n, err := reader.Read(buf[:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if n > 0 {
			offset++
			if buf[0] == '\n' {
				line--
			}
		}
	}
	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}
	return nil
}

func (_this *CSV) lines() error {
	file, err := _this.Open()
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_this.SeekToLine(file, 4)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		fmt.Println(reader.Text())
	}
	return nil
}

func (_this *CSV) Open() (*os.File, error) {
	file, err := os.Open(_this.pathFile)
	if err != nil {
		return nil, ErrOpenFile
	}
	return file, nil
}

func (_this *CSV) yield() (<-chan CSVLine, error) {
	file, err := _this.Open()
	if err != nil {
		return nil, err
	}
	line := NewCSVLine(_this._headers)
	_this.SeekToLine(file, 1)
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	chnl := make(chan CSVLine)
	go func() {
		for scanner.Scan() {
			line.parse(scanner.Text())
			chnl <- *line
		}
		close(chnl)
	}()
	return chnl, nil
}

func (_this *CSV) parseHeader() error {
	file, err := _this.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	header, _, err := reader.ReadLine()
	if err != nil {
		return ErrCouldNotReadTheFile
	}
	headers := bytes.Split(header, string(_this.delimiter))
	for i, header := range headers {
		_this._headers[string(header)] = i
	}
	_this.currPos++
	return nil
}



func (_this *CSV) parseDelimeter() error {
	file, err := _this.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	header, _, err := reader.ReadLine()
	if err != nil {
		return ErrCouldNotReadTheFile
	}

	headers := bytes.Split(header, string(_this.delimiter))
	for i, header := range headers {
		_this._headers[string(header)] = i
	}
	_this.currPos++
	return nil
}
