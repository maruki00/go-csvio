package consts

import "errors"

var (
	ErrOpenFile            = errors.New("could not open the csv file")
	ErrFileNotAccessible   = errors.New("could not access this csv file")
	ErrFileProbablyEmpty   = errors.New("file probably empty")
	ErrCouldNotReadTheFile = errors.New("could not read the file")
)
