package anansi

import (
	"fmt"
	"log"
	"time"
)

func init() {
	log.Println("init in util.go")
}

// tableFormat defines cli table formatting options
type tableFormat struct {
	TopHeader       string
	BottomHeader    string
	MiddleHeader    string
	SideBorder      string
	DebugHeader     string
	AlternateBorder bool
	AlternateColor  bool
	HeaderColor     ansiCodes
	RowColors       []ansiCodes
	BorderColor     ansiCodes
}

// codeSymbolColors defines the ANSI colors used for 'pretty' code output
type codeSymbolColors struct {
	Keywords    Attribute
	Operator    Attribute
	Integers    Attribute
	Floats      Attribute
	Strings     Attribute
	Comments    Attribute
	Punctuation Attribute
	Types       Attribute
}

type defaultValues struct {
	ScreenWidth      int
	UseColor         bool
	TableFormat      tableFormat
	CodeSymbolColors codeSymbolColors
}

var config Session

// VerboseLevel defines the level of visual feedback in the cli terminal
type VerboseLevel uint

const (
	verboseQuiet VerboseLevel = iota
	verboseMinimal
	verboseStandard
	verboseVerbose
	verboseDebug
	verboseAll
)

// StartTimer returns a function that will return the elapsed time
func StartTimer(name string) func() {
	/*
		    func example() {
			   stop := StartTimer("myTimer")
			   defer stop()

			   time.Sleep(1*time.Second)
	*/

	t := time.Now()
	logPrint("Timer started:", name)
	return func() {
		d := time.Now().Sub(t)
		logPrint(name, "took", d)
	}
}

func br() (n int, err error) {
	return fmt.Println("")
}

func hr() (n int, err error) {
	return fmt.Println("")
}

// Timestamp code license information located at /Users/xxxxxx/go/pkg/mod/google.golang.org/protobuf@v1.23.0/types/known/timestamppb/timestamp.pb.go in go 1.15.5 darwin/amd64
//
// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// https://developers.google.com/protocol-buffers/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
