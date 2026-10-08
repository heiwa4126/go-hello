package main

import (
	"fmt"
	"runtime/debug"

	"github.com/heiwa4126/go-hello/say"
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

	message := say.Say("Hello, World!")
	fmt.Println(message)
}
