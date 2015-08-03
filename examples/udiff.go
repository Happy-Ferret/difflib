package main

import (
	"fmt"
	"strings"

	".."
)

const (
	fileA = `
1
2
3
4
this line was removed
5
6
7
8
`
	fileB = `
1
2
3
4
this is an added line
5
6
7
8
`
)

func main() {

	fmt.Println(difflib.UDiff("fileA", "fileB", strings.Split(fileA, "\n"), strings.Split(fileB, "\n")))

}
