// package gocsvio
package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

var (
	ErrOpenFile            = errors.New("could not open the csv file")
	ErrFileNotAccessible   = errors.New("could not access this csv file")
	ErrFileProbablyEmpty   = errors.New("file probably empty")
	ErrCouldNotReadTheFile = errors.New("could not read the file")
)

type CSV struct {
	_headers   map[string]int
	delimiters []byte
	file       *os.File
	currPos    int64
}

func (_this *CSV) SetDelimiters(delimiters []byte) {
	_this.delimiters = delimiters
}

func (_this *CSV) SeekToLine( /* file *os.File, */ line uint) error {
	var buf [1]byte
	var offset = int64(1)
	_this.file.Seek(int64(0), io.SeekStart)
	reader := bufio.NewReader(_this.file)
	for _, err := reader.Read(buf[:]); err != nil && line > 0; {
		if buf[0] == '\n' {
			line--
		}
		_this.file.Seek(offset, io.SeekCurrent)
	}
	return nil
}

func NewReader(csvpath string, defaultSep []byte) (*CSV, error) {
	file, err := os.Open(csvpath)
	if err != nil {
		return nil, ErrOpenFile
	}
	obj := &CSV{
		_headers:   make(map[string]int),
		delimiters: defaultSep,
		file:       file,
		currPos:    int64(0),
	}
	if err := obj.parseHeader(); err != nil {
		return nil, err
	}

	return obj, nil
}

func (_this *CSV) lines() (<-chan string, error) {

	// //test
	// file, err := os.Open("./file.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// scanner := bufio.NewScanner(file)
	// file.Seek(int64(1), 1)
	//
	// // end test

	_this.file.Seek(int64(1), io.SeekCurrent)
	scanner := bufio.NewScanner(_this.file)
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	chnl := make(chan string)
	go func() {
		for scanner.Scan() {
			chnl <- scanner.Text()
		}
		close(chnl)
	}()
	return chnl, nil
}

func (_this *CSV) parseHeader() error {
	if _this.file == nil {
		return ErrFileNotAccessible
	}
	reader := bufio.NewReader(_this.file)
	// _, err := _this.file.Seek(int64(0), 0)
	// if err != nil {
	// 	return ErrFileProbablyEmpty
	// }
	header, _, err := reader.ReadLine()
	if err != nil {
		return ErrCouldNotReadTheFile
	}
	headers := bytes.Split(header, _this.delimiters)
	for i, header := range headers {
		_this._headers[string(header)] = i
	}
	_this.currPos++
	return nil
}

func (_this *CSV) Get(key string) any {
	index, ok := _this._headers[key]
	if !ok {
		return ""
	}
	return index
}
