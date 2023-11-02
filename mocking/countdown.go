package main

import (
	"fmt"
	"io"
)

func main() {
}

func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}
