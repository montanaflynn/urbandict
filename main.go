package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug = kingpin.Flag("debug", "Enable debug mode.").Short('d').Bool()
	synonyms  = kingpin.Flag("synonyms", "Retrieve synonyms instead of definition").Short('s').Bool()
	word  = kingpin.Arg("word", "Word to do.").Required().String()
)

func main() {
	kingpin.New("urbandict", "Urban Dictionary from the command line.")
	kingpin.Version("0.0.1")
	kingpin.Parse()
	startUrbanDict()
}
