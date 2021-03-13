package main

import (
	"github.com/rwxrob/cmdtab"
	"github.com/rwxrob/uniq-go"
)

func main() {
	cmdtab.SmartPrintln(uniq.Base32())
}
