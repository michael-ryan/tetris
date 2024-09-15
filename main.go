package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/michael-ryan/tetris/game"
)

func main() {
	isGui := flag.Bool("gui", false, "Use GUI?")
	flag.Parse()
	if *isGui {
		gui()
	} else {
		cli()
	}
}

func gui() {
	panic("gui not implemented")
}

func cli() {
	game, err := game.NewGame()
	if err != nil {
		fmt.Printf("error loading game: %v", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		reader.ReadBytes('\n')
		for i := 0; i < 100; i++ {
			fmt.Println()
		}
		game.Step()
		fmt.Println(game.Draw())
	}
}
