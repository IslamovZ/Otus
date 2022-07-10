package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	myReversedString := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(myReversedString)
}
