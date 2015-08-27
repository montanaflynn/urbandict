package main

import (
	"fmt"
	"github.com/pkg/term"
	"os"
)

const (
	HELP_MSG = "(Use arrow keys to cycle and q to quit)"
)

var (
	current = 0
	count   = 0
)

func startUrbanDict() {
	def, err := getDefinitions(*word)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *tags == true {
		count = len(def.Tags)
	} else {
		count = len(def.List)
	}

	display(def, current)

	for {
		handleInput(def)
	}

}

func display(d data, i int) {
	if *tags == true {
		displayTag(d, i)
	} else {
		displayDefinition(d, i)
	}
}

func displayTag(d data, i int) {
	clearTerm()
	tagTemplate := "Tag #%v for %s %s\n\n"
	fmt.Printf(tagTemplate, i+1, *word, HELP_MSG)
	fmt.Printf("%s\n", d.Tags[i])
}

func displayDefinition(d data, i int) {
	di := d.List[i]
	clearTerm()
	defineTemplate := "Definition #%v for %s +%v -%v %s\n\n"
	fmt.Printf(defineTemplate, i+1, di.Word, di.ThumbsUp, di.ThumbsDown, HELP_MSG)
	fmt.Printf("%s\n", di.Definition)
}

func handleInput(def data) {
	ascii, keyCode, err := getChar()
	if err != nil {
		fmt.Println("Unexpected error retriving character code")
		os.Exit(0)
	}
	if keyCode == 37 || keyCode == 38 {
		if current > 0 {
			current -= 1
		} else {
			current = count - 1
		}
	} else if keyCode == 39 || keyCode == 40 {
		if current < count-1 {
			current += 1
		} else {
			current = 0
		}
	} else if ascii == 113 {
		os.Exit(0)
	}

	display(def, current)

}

// Find the ascii code or if input is an arrow a Javascript key code
func getChar() (ascii int, keyCode int, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
		if bytes[2] == 65 {
			// Up
			keyCode = 38
		} else if bytes[2] == 66 {
			// Down
			keyCode = 40
		} else if bytes[2] == 67 {
			// Right
			keyCode = 39
		} else if bytes[2] == 68 {
			// Left
			keyCode = 37
		}
	} else if numRead == 1 {
		ascii = int(bytes[0])
	}
	t.Restore()
	t.Close()
	return
}

func clearTerm() {
	print("\033[H\033[2J")
}
