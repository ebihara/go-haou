package main

import (
	"fmt"
	"os/user"
)

type HaouGuesser struct {}
func (p *HaouGuesser) isHaou() bool {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	blacklist := []string{"co3k"}
	for i :=0; i < len(blacklist); i++ {
		if (usr.Username == blacklist[i]) {
			return false;
		}
	}

	return true
}

func main() {
	guesser := HaouGuesser{}

	if (guesser.isHaou()) {
		fmt.Println("私が覇王です")
	} else {
		fmt.Println("私は覇王ではないです")
	}
}
