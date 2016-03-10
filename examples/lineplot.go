package main

import (
	"flag"
	"os"
)

const Version = "0.0"

func main() {

	version := flag.Bool("version", false, "Print version number")
	flag.Parse()

	if *version {
		println(Version)
		os.Exit(0)
	}

}

