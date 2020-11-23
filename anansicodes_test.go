package anansi

import (
	"testing"
)

//* --------------------------------------------------------> Attribute type definition
func TestAttribute_Ansi(t *testing.T) {
	tests := []struct {
		name string
		a    Attribute
		want string
	}{
		// TODO: Add test cases.
		{"Test Ansi Black", FgBlack, "\x1b[30m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Ansi(); got != tt.want {
				t.Errorf("Attribute.Ansi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttribute_Bright(t *testing.T) {
	tests := []struct {
		name string
		a    Attribute
		want string
	}{
		// TODO: Add test cases.
		{"Test Bright Black", FgBlue, "\x1b[34;1m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Bright(); got != tt.want {
				t.Errorf("Attribute.Bright() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttribute_A256(t *testing.T) {
	type args struct {
		fb string
	}
	tests := []struct {
		name string
		a    Attribute
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test Ansi 134", 134, args{fg}, "\x1b[38;5;134m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.A256(tt.args.fb); got != tt.want {
				t.Errorf("Attribute.A256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttribute_Fg(t *testing.T) {
	tests := []struct {
		name string
		a    Attribute
		want string
	}{
		// TODO: Add test cases.
		{"Test Ansi 134", 134, "\x1b[38;5;134m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Fg(); got != tt.want {
				t.Errorf("Attribute.Fg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttribute_Bg(t *testing.T) {
	tests := []struct {
		name string
		a    Attribute
		want string
	}{
		// TODO: Add test cases.
		{"Test Ansi 134", 134, "\x1b[48;5;134m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Bg(); got != tt.want {
				t.Errorf("Attribute.Bg() = %v, want %v", got, tt.want)
			}
		})
	}
}

//* --------------------------------------------------------> ansiCodes type definition
func Test_ansiCodes_Set(t *testing.T) {
	type fields struct {
		fg Attribute
		bg Attribute
		ef Attribute
	}
	type args struct {
		fg Attribute
		bg Attribute
		ef Attribute
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ansiCodes{
				fg: tt.fields.fg,
				bg: tt.fields.bg,
				ef: tt.fields.ef,
			}
			a.Set(tt.args.fg, tt.args.bg, tt.args.ef)
		})
	}
}

func Test_ansiCodes_Reset(t *testing.T) {
	type fields struct {
		fg Attribute
		bg Attribute
		ef Attribute
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ansiCodes{
				fg: tt.fields.fg,
				bg: tt.fields.bg,
				ef: tt.fields.ef,
			}
			if got := a.Reset(); got != tt.want {
				t.Errorf("ansiCodes.Reset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ansiCodes_Print(t *testing.T) {
	type fields struct {
		fg Attribute
		bg Attribute
		ef Attribute
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ansiCodes{
				fg: tt.fields.fg,
				bg: tt.fields.bg,
				ef: tt.fields.ef,
			}
			a.Print()
		})
	}
}

func Test_ansiCodes_String(t *testing.T) {
	type fields struct {
		fg Attribute
		bg Attribute
		ef Attribute
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ansiCodes{
				fg: tt.fields.fg,
				bg: tt.fields.bg,
				ef: tt.fields.ef,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("ansiCodes.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ansiCodes_Wrap(t *testing.T) {
	type fields struct {
		fg Attribute
		bg Attribute
		ef Attribute
	}
	type args struct {
		s []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ansiCodes{
				fg: tt.fields.fg,
				bg: tt.fields.bg,
				ef: tt.fields.ef,
			}
			if got := a.Wrap(tt.args.s...); got != tt.want {
				t.Errorf("ansiCodes.Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
