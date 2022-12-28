package main

import (
	"github.com/getlantern/systray"
	interna "github.com/lw396/variable/interna"
)

func main() {
	onExit := func() {}
	systray.Run(interna.OnReady, onExit)
}
