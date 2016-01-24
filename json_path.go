package main

import "flag"

type JSONPath *string

func NewJSONPath() JSONPath {
	var (
		p = new(string)

		defaultValue = ""

		option      = "json-path"
		shortOption = "j"

		usage      = "JSON path to verify"
		shortUsage = usage + " (shorthand)"
	)

	flag.StringVar(p, option, defaultValue, usage)
	flag.StringVar(p, shortOption, defaultValue, shortUsage)

	return p
}
