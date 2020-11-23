package anansi

import "strings"

// Anansi - ANSI color for terminal output
// Copyright (c) 2020 Michael Treanor
// MIT License

//* --------------------------------------------------------> ansiCodes type definition

// AString  represents an ANSI color formatted string builder.
type AString struct {
	sb     *strings.Builder
	sbAnsi ansiCodes
}

// Fg adds a foreground ANSI 256 color code to the string
func (a *AString) Fg(code Attribute) {
	a.sb.WriteString(code.Fg())
}

// Bg adds a background ANSI 256 color code to the string
func (a *AString) Bg(code Attribute) {
	a.sb.WriteString(code.Bg())
}

// Ef adds an ANSI code to the string
func (a *AString) Ef(code Attribute) {
	a.sb.WriteString(code.Ansi())
}

// Reset resets all ANSI escape codes
func (a *AString) Reset() {
	a.sb.WriteString(Reset.Ansi())
}

// Clear resets the string to use zero bytes so that it may be garbage-collected.
func (a *AString) Clear() {
	a.sb.Reset()
}

// String returns the complete string buffer
func (a *AString) String() string {
	return a.sb.String()
}
