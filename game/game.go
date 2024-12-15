package game

import (
	"DungeonGame/params"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player  Player
	enemies Enemies
}

func NewGame() *Game {
	game := &Game{}

	game.player = *NewPlayer(game)
	game.enemies = *SpawnEnemies(game)

	return game
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawMap(screen)
	g.player.Draw(screen) // Draw the player
	g.enemies.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return params.ScreenWidth, params.ScreenHeight
}
