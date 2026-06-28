// Package main is a minimal, dependency-free Go program used as a curated
// build-from-source sample for Stout end-to-end validation.
package main

import "fmt"

// Greeting returns the canonical greeting. Kept as an exported function so the
// build produces a non-trivial, testable unit.
func Greeting() string {
	return "hello from a Stout curated build"
}

func main() {
	fmt.Println(Greeting())
}
