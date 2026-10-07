package main

import (
	"fmt"
	"runtime/debug"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

var (
	Version  = ""
	Revision = ""
)

func init() {
	info, ok := debug.ReadBuildInfo()
	if ok {
		if Version == "" {
			Version = info.Main.Version
		}
		if Revision == "" {
			for _, setting := range info.Settings {
				if setting.Key == "vcs.revision" {
					Revision = setting.Value
					break
				}
			}
		}
	}

	if Version == "" {
		Version = "(devel)"
	}
	if Revision == "" {
		Revision = "(unknown)"
	}
}

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
