package main

// func type
type LoggerAdapter func(message string)

// function method -- takes incompatible "LogOutput" function
// implementation and makes it satisfy the `Logger` interface
func (lg LoggerAdapter) Log(message string) {
	lg(message)
}
