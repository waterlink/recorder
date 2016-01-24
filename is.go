package main

import (
	"errors"
	"flag"
	"fmt"
)

type Is struct {
	args     []string
	index    Index
	jsonPath JSONPath
}

func IsAction(args []string) error {
	return (&Is{
		args:     args,
		index:    NewIndex(),
		jsonPath: NewJSONPath(),
	}).Run()
}

func (i *Is) Run() error {
	flag.Usage = UsageFor("is METHOD PATH = DATA")
	flag.CommandLine.Parse(i.args)

	args := flag.Args()
	if len(args) != 4 || args[2] != "=" {
		flag.Usage()
		return errors.New("Expected 4 arguments: METHOD PATH = DATA")
	}

	method, path, data := args[0], args[1], args[3]

	record, err := LoadRecord(method, path, *i.index)
	if err != nil {
		return fmt.Errorf("Unable to load record: %s", err)
	}

	if record.Data != data {
		return fmt.Errorf("Expectation failed on record's data:\n\tExpected: %s\n\tActual:   %s\n\n", data, record.Data)
	}

	return nil
}
