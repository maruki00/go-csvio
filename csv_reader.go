// package gocsvio
package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)



var delimiters = []byte{',', '|', ';', '\t', '`', '"', '~'}




