package anansi

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

// Session defines session information
//
// Part of the data is a Timer that returns start time (timestamp) and a function that will return a time interval (from the start time) when called
//
// Common Timestamp functions
    //
    //     type Timestamp
    //     func New(t time.Time) *Timestamp
    //     func Now() *Timestamp
    //     func (x *Timestamp) AsTime() time.Time
    //     func (x *Timestamp) CheckValid() error
    //     func (x *Timestamp) GetNanos() int32
    //     func (x *Timestamp) GetSeconds() int64
    //     func (x *Timestamp) IsValid() bool
    //     func (x *Timestamp) Reset()
    //     func (x *Timestamp) String() string
type Session interface {
    Name() string
    IsDevMode() bool
    UseLogger() bool
    Verbose() VerboseLevel
    temp() *os.File
    SessionStart() *time.Time
    SessionEnd() *time.Time
    UserID() uint
}

type VerboseLevel uint
const (
    Quiet        VerboseLevel = iota
    Minimal
    Standard
    Verbose
    Debug
    All
)

func  (s *Session) SessionStart() *time.Time {
    return time.Now()
}


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


func logPrint(v ...interface{}) {
	if config.VerboseLevel < debug.DEV_OUTPUT 0 {

		log.Println("----------")
		defer log.Println("----------")
		log.Println(v)
	}
	// todo something here ...
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
