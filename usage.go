package main

import (
	"flag"
	"fmt"
	"os"
)

func UsageFor(action string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "Usage of %s %s:\n", os.Args[0], action)
		flag.PrintDefaults()
	}
}
