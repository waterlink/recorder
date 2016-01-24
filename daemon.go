package main

import "fmt"

func DaemonAction(args []string) error {
	fmt.Printf("Daemon will start with args: %#v\n", args)
	return nil
}
