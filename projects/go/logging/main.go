//+build !test

package main

import (
	"os"
)

func main() {
	os.Exit(exec(os.Args))
}
