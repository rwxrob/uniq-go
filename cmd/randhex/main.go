package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rwxrob/cmdtab"
	"github.com/rwxrob/uniq-go"
)

func main() {
	n := 18
	var err error
	if len(os.Args) > 1 {
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			os.Exit(1)
		}
	}
	t := fmt.Sprintf("%x", uniq.Bytes(n))
	cmdtab.SmartPrintln(t)
}
