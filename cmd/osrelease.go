package main

import (
	"fmt"

	"github.com/stobbsm/go-osrelease"
)

// Main executable for the osrelease library. Should help with using
// different os release values in scripts

func usage() {
	// Print usage statement for users
}

func main() {
	fmt.Print("go-osrelease")

	usage()

	i := osrelease.NewOsInfo().Load()

	fmt.Print(i.Id())
}
