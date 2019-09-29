package main

import (
	"fmt"
	flag "github.com/ogier/pflag"
	"os"
)

const version = "0.0.0"

func main() {
	fmt.Println("Loading Gordon IRC bot...")
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n\n")
		flag.PrintDefaults()
	}

	showVersionInfo := flag.BoolP("version", "V", false, "Print version info and quit.")
	flag.Parse()

	if *showVersionInfo {
		fmt.Println("Gordon " + version)
		os.Exit(0)
	}
}
