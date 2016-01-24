package main

import "os"

func ResetAction(args []string) error {
	dir, err := BaseDir()
	if err != nil {
		return err
	}

	return os.RemoveAll(dir)
}
