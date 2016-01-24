package main

import "os"

var (
	actions = map[string]Action{
		"daemon": DaemonAction,
		"get":    GetAction,
		"reset":  ResetAction,
		"is":     IsAction,
	}
)

func main() {
	if len(os.Args) < 2 {
		Fatal("Expected an action\n")
	}

	actionName := os.Args[1]
	action, ok := actions[actionName]
	if !ok {
		Fatalf("Action '%s' is unknown\n", actionName)
	}

	if err := action(os.Args[2:]); err != nil {
		Fatalf("Encountered an error: %s\n", err)
	}
}
