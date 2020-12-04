package main

import (
	"io"
	"os"

	"github.com/skeptycal/anansi"
)

func main() {
	io.Copy(anansi.Output, os.Stdin)
}
