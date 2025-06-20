package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pkg/profile"
)

type CSV struct {
	headers map[string]int
	csvPath string
	pos     int64
	sep     byte
}

func NewCSV(csvPath string, sep byte) *CSV {
	obj := &CSV{
		sep:     sep,
		csvPath: csvPath,
		pos:     0,
		headers: make(map[string]int),
	}
	err := obj.ReadHeaders()
	if err != nil {
		fmt.Printf("failed to read headers: %v\n", err)
	}
	return obj
}

func (_this *CSV) SetHeaders(headers map[string]int) {
	_this.headers = headers
}

func (_this *CSV) ReadHeaders() error {
	f, err := os.OpenFile(_this.csvPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Seek(int64(0), io.SeekStart)

	reader := bufio.NewReader(f)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return err
	}
	line = strings.TrimRight(line, "\r\n")
	headers := strings.Split(line, string(_this.sep))
	for i, header := range headers {
		_this.headers[header] = i
	}
	_this.pos = int64(len(line) + 1)
	return nil
}

func (_this *CSV) ReadLine() (<-chan [][]byte, error) {
	file, err := os.Open(_this.csvPath)
	if err != nil {
		return nil, err
	}
	if _, err := file.Seek(_this.pos, io.SeekStart); err != nil {
		file.Close()
		return nil, err
	}
	stream := make(chan [][]byte, 64)
	go func() {
		defer file.Close()
		defer close(stream)
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if err != io.EOF {
					log.Println(err.Error())
				}
				break
			}
			line = bytes.TrimRight(line, "\r\n")
			stream <- bytes.Split(line, []byte{_this.sep})
		}
	}()
	return stream, nil
}

func (_this *CSV) Write(filename string, lines <-chan [][]byte) error {
	var err error
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	if _, err := writer.WriteString(strings.Join(_this.headersValues(), string(_this.sep)) + "\n"); err != nil {
		return fmt.Errorf("error writing header: %w", err)
	}
	for line := range lines {
		_, err = writer.Write(append(bytes.Join(line, []byte{_this.sep}), '\n'))
		if err != nil {
			// return fmt.Errorf("error writing record: %w", err)
		}
	}
	return nil
}

func (_this *CSV) Get(key string, line [][]byte) []byte {
	if index, ok := _this.headers[key]; ok {
		return line[index]
	}
	return []byte{}
}

func (_this *CSV) headersValues() []string {
	headers := make([]string, len(_this.headers))
	for name, idx := range _this.headers {
		if idx >= 0 && idx < len(headers) {
			headers[idx] = string(name)
		}
	}
	return headers
}

func main() {

	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	csv := NewCSV("./file.csv", ',')
	lines, err := csv.ReadLine()
	if err != nil {
		panic(err)
	}

	for line := range lines {

		fmt.Println(line)
	}

	// file, err := os.Open("file.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// parts := make([]string, 0)
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	_ = Split(scanner.Text(), ",", &parts)
	// }
}
