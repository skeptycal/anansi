package anansi

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	db "github.com/skeptycal/godebug"
)

func init() {
	log.Println("anansi logging started")
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

var config db.Session

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
	log.Info("Timer started:", name)
	return func() {
		d := time.Now().Sub(t)
		log.Info(name, "took", d)
	}
}

func br() (n int, err error) {
	return fmt.Println("")
}

func hr() (n int, err error) {
	return fmt.Println("")
}
