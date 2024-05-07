package main

import (
	"fmt"

	"github.com/kiritocyanpine/grpctest/sample"
)

func main() {
	l := sample.NewLaptop()
	fmt.Println("Hello", l)
}
