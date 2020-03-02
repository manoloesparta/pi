package search

import (
	"io/ioutil"
	"strconv"
	"strings"
)

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
	for i := -128; i < 128; i++ {
		if i == 0 {
			sb.WriteString("s")
		} else if i == digits {
			sb.WriteString("e")
		}
		sb.WriteByte((*content)[index+i])
	}
	return sb.String()
}
