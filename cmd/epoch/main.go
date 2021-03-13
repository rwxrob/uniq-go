package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rwxrob/cmdtab"
	"github.com/rwxrob/uniq-go"
)

func main() {
	if len(os.Args) <= 1 {
		cmdtab.SmartPrintln(uniq.Second())
		os.Exit(1)
	}
	i, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	t := fmt.Sprintf("%v", time.Unix(i, 0))
	cmdtab.SmartPrintln(t)
}
