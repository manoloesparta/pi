package search

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const chunkSize int = 128

// Get get the string that will be displayed
func Get(input string) (string, int) {
	index, digits := allResults(input)
	fileNum := 0

	if index > digitsOfFile {
		fileNum = index / digitsOfFile
		index = index - (fileNum * digitsOfFile)
	}

	content := contentOfFile(fileNum)
	res := makeString(digits, index, &content)

	return res, index
}

func contentOfFile(fileNum int) string {
	f := "num/part" + strconv.Itoa(fileNum)
	file, _ := ioutil.ReadFile(f)
	return string(file)
}

func makeString(digits int, index int, content *string) string {
	var sb strings.Builder
	start, end := getBounds(index)
	fmt.Println(start, end)
	for i := start; i < end; i++ {
		if i == 0 {
			sb.WriteByte('s')
		} else if i == digits {
			sb.WriteByte('e')
		}
		sb.WriteByte(' ')
		sb.WriteByte((*content)[index+i])
	}
	return sb.String()
}

func getBounds(index int) (int, int) {
	res := index - chunkSize
	if res < 0 {
		return 0, chunkSize * 2
	}
	return -chunkSize, chunkSize
}
