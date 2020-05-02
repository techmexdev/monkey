package main

import (
	"log"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Printf("failed reading current os user = %+v\n", err)
	}

	log.Printf("Hello %s. Welcome to the Monkey REPL!", usr.Username)

	repl.Start(os.Stdin, os.Stdout)
}
