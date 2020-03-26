package main

import (
	"fmt"
	"os"
	"os/user"
	"rainbow/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Rainbow programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
