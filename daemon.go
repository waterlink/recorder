package main

import "flag"

type Daemon struct {
	args   []string
	listen Listen
}

func DaemonAction(args []string) error {
	return (&Daemon{
		args:   args,
		listen: NewListen(),
	}).Start()
}

func (d *Daemon) Start() error {
	flag.Usage = UsageFor("daemon")
	flag.CommandLine.Parse(d.args)

	return NewServer(d.listen).Start()
}
