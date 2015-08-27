package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug = kingpin.Flag("debug", "Enable debug mode.").Short('d').Bool()
	tags  = kingpin.Flag("tags", "Retrieve tags instead of definition").Short('t').Bool()
	word  = kingpin.Arg("word", "Word to do.").Required().String()
)

func main() {
	kingpin.New("urbandict", "Urban Dictionary from the command line.")
	kingpin.Version("0.0.1")
	kingpin.Parse()
	startUrbanDict()
}
