package main

import (
	"os"
	"path/filepath"

	"github.com/goose-play/gocli"
)

func main() {
	boot := gocli.CLI()
	args, fmap := gocli.ParseInputArgs(os.Args[1:])
	path, _ := filepath.Abs(".")
	if _, ok := fmap.HasFlag(gocli.WorkDir); !ok {
		fmap.Set(gocli.WorkDir.Name(), path)
	}
	msg := boot.Run(gocli.MergeFlagMap(args, fmap))
	os.Stdout.WriteString(msg.Msg())
}
