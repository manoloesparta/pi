package main

import (
	"fmt"

	"github.com/maybemanolo/pi/api/search"
)

func main() {
	index, length := search.AllResults("663270532")
	fmt.Println(index, length)
}
