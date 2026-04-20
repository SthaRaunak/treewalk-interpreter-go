package main

import (
	"fmt"
	"os"
	"os/user"

	"raunak.sh/repl"
)

func main() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Print("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
