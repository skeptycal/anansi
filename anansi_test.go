package anansi

// Anansi - ANSI color for terminal output
// Copyright (c) 2020 Michael Treanor
// MIT License

import (
	"testing"
)

func TestAnsi(t *testing.T) {
	type args struct {
		color Attribute
	}
	tests := []struct {
		name  string
		color Attribute
		want  string
	}{
		{"FgBlack color test", FgBlack, "\x1b[30m"},
		{"FgRed color test", FgRed, "\x1b[31m"},
		{"FgGreen color test", FgGreen, "\x1b[32m"},
		{"FgYellow color test", FgYellow, "\x1b[33m"},
		{"FgBlue color test", FgBlue, "\x1b[34m"},
		{"FgMagenta color", FgMagenta, "\x1b[35m"},
		{"FgCyan color test", FgCyan, "\x1b[36m"},
		{"FgWhite color test", FgWhite, "\x1b[37m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ansi(tt.color); got != tt.want {
				t.Errorf("Ansi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllAnsi(t *testing.T) {
	type args struct {
		fg Attribute
		bg Attribute
		ef Attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Normal white on black", args{FgWhite, BgBlack, Reset}, "[37m;[40m;[0m"},
		{"Bold Red on Yellow", args{FgRed, BgYellow, Bold}, "[31m;[43m;[1m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllAnsi(tt.args.fg, tt.args.bg, tt.args.ef); got != tt.want {
				t.Errorf("AllAnsi() = %v, want %v", got, tt.want)
			}
		})
	}
}
