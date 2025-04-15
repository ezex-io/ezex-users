// Package main defines the entry point for the program.
package main

import "log"

// Greet returns a greeting message.
func Greet() string {
	return "Hello, world!"
}

func main() {
	log.Println(Greet())
}
