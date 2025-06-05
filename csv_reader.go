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
	_headers  map[string]int
	line      []string
	seperator []byte
	fd        *os.File
	currPos   int64
}

func NewReader(csvpath string, defaultSep []byte) (*CSV, error) {
	fd, err := os.Open(csvpath)
	if err != nil {
		return nil, ErrOpenFile
	}
	obj := &CSV{
		_headers:  make(map[string]int),
		line:      make([]string, 0),
		seperator: defaultSep,
		fd:        fd,
		currPos:   int64(0),
	}
	if err := obj.parseHeader(); err != nil {
		return nil, err
	}
	return obj, nil
}

func (_this *CSV) lines() (<-chan string, error) {
	_this.fd.Seek(int64(1), io.SeekCurrent)
	scanner := bufio.NewScanner(_this.fd)
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
	if _this.fd == nil {
		return ErrFileNotAccessible
	}
	reader := bufio.NewReader(_this.fd)
	_, err := _this.fd.Seek(int64(0), 0)
	if err != nil {
		return ErrFileProbablyEmpty
	}
	header, _, err := reader.ReadLine()
	if err != nil {
		return ErrCouldNotReadTheFile
	}
	headers := bytes.Split(header, _this.seperator)
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
	return _this.line[index]
}
