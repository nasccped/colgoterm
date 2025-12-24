package app

import "github.com/nasccped/colgoterm/internals/colored"

// Stores function pointer to generate background color.
type functionPointer struct {
	// For normal color background.
	normal func(*colored.Colored) *colored.Colored
	// For bright color background.
	bright func(*colored.Colored) *colored.Colored
}

// Creates a new `functionPointer` struct.
func newFunctionPointer(normal, bright func(*colored.Colored) *colored.Colored) *functionPointer {
	return &functionPointer{
		normal: normal,
		bright: bright,
	}
}
