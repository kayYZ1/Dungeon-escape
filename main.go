package main

import (
	"DungeonGame/game"
	"DungeonGame/params"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowTitle("Dungeon Escape")
	ebiten.SetWindowSize(params.ScreenWidth*2, params.ScreenHeight*2)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
