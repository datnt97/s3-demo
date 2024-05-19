package core

import "runtime/debug"

func PrintStack() {
	debug.PrintStack()
}

func SprintStack() string {
	return string(debug.Stack())
}
