package main

import "fmt"

var VERSION string

func VersionAction(args []string) error {
	version := VERSION
	if version == "" {
		version = "development"
	}

	fmt.Printf("recorder %s\n", version)
	return nil
}
