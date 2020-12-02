package main

import (
	"io"
	"os"

	"github.com/skeptycal/util/anansi"
)

func main() {
	io.Copy(anansi.Output, os.Stdin)
}
