package main

import (
	"errors"
	"flag"
	"fmt"
)

type Get struct {
	args  []string
	index Index
}

func GetAction(args []string) error {
	return (&Get{
		args:  args,
		index: NewIndex(),
	}).Run()
}

func (g *Get) Run() error {
	flag.Usage = UsageFor("get METHOD PATH")
	flag.CommandLine.Parse(g.args)

	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		return errors.New("Expecting 2 arguments: METHOD PATH")
	}

	method, path := args[0], args[1]

	record, err := LoadRecord(method, path, *g.index)
	if err != nil {
		return fmt.Errorf("Unable to load record: %s", err)
	}

	fmt.Println(record.Data)
	return nil
}
