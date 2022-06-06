package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const stringForReverse = "Hello, OTUS!"

func printReverseString(str string) {
	myReversedString := stringutil.Reverse(str)
	fmt.Println(myReversedString)
}

func main() {
	printReverseString(stringForReverse)
}
