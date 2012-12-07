package main

import (
	"fmt"
)

type HaouGuesser struct {}
func (p *HaouGuesser) isHaou() bool { return true }

func main() {
	guesser := HaouGuesser{}

	if (guesser.isHaou()) {
		fmt.Println("私が覇王です")
	} else {
		fmt.Println("私は覇王ではないです")
	}
}
