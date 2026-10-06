package main

import (
	"fmt"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

var (
	Version  = "dev"
	Revision = "unknown"
)

func main() {

	fmt.Printf("Version: %s\nRevision: %s\n", Version, Revision)

	say, err := cowsay.Say(
		"Hello, World!",
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
}
