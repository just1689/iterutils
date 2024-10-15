package main

import (
	"fmt"
	"github.com/just1689/iterutils/iterutils"
	"strings"
)

func main() {
	count := []string{"one", "two", "three"}
	list := iterutils.
		From(count).
		Filter(filterOutNoEs).
		Map(toUpper).
		Reverse().
		Collect()
	fmt.Println(list)
	listSize := iterutils.MapTo(list, mapper)
	fmt.Println(listSize)

}

func filterOutNoEs(s string) bool {
	return strings.Contains(s, "e")
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func mapper(in string) int {
	return len(in)
}
