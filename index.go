package main

import "flag"

const LastIndex = -1

type Index *int

func NewIndex() Index {
	var (
		i = new(int)

		defaultValue = LastIndex

		option      = "index"
		shortOption = "i"

		usage      = "HTTP request index to operate upon"
		shortUsage = usage + " (shorthand)"
	)

	flag.IntVar(i, option, defaultValue, usage)
	flag.IntVar(i, shortOption, defaultValue, shortUsage)

	return i
}
