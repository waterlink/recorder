package main

import "flag"

type Listen *string

func NewListen() Listen {
	var (
		l = new(string)

		defaultValue = ":9977"

		option      = "listen"
		shortOption = "l"

		usage      = "HTTP server listen binding"
		shortUsage = usage + " (shorthand)"
	)

	flag.StringVar(l, option, defaultValue, usage)
	flag.StringVar(l, shortOption, defaultValue, shortUsage)

	return l
}
