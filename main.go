package main

import (
	"github.com/omerkaya1/notification/cmd"
	"log"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
