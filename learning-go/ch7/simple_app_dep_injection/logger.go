package main

// required Logger interface (not compatible with LogOutput function)
type Logger interface {
	Log(message string)
}
