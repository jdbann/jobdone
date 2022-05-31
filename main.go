package main

import (
	"fmt"
	"os"

	"jobdone.emailaddress.horse/cmd"
)

func main() {
	err := cmd.Start(os.Args[1:])
	if err != nil {
		fmt.Printf("OH NO! There has been an error: %v", err)
		os.Exit(1)
	}
}
