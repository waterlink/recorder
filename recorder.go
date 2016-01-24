package main

import (
	"fmt"
	"os"
)

type Action func(args []string) error

func Fatal(s string) {
	fmt.Print(s)
	os.Exit(1)
}

func Fatalf(s string, args ...interface{}) {
	fmt.Printf(s, args...)
	os.Exit(1)
}
