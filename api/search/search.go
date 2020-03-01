package search

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

const digitsOfFile int = 100000000

// AllResults performs a fullScan ofn every word and if it doesnt find it reduces the string
func AllResults(input string) (int, int) {
	long := len(input)
	for i := 0; i < long; i++ {
		res, err := fullScan(input[0 : long-i])
		if err == nil {
			return res, long - i
		}
	}
	return -1, -1
}

func singleSearch(input string, fileindex int) (int, error) {
	f := "../num/part" + strconv.Itoa(fileindex)
	file, _ := ioutil.ReadFile(f)
	content := string(file)
	index := strings.Index(content, input)

	if index == -1 {
		return -1, errors.New("Not found")
	}

	return index, nil
}

func fullScan(input string) (int, error) {
	for i := 0; i < 10; i++ {
		res, err := singleSearch(input, i)
		if err == nil {
			return res + (digitsOfFile * i), nil
		}
	}
	return -1, errors.New("Non existent")
}
