package game

import (
	"DungeonGame/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	X      int
	Y      int
	damage int
	model  *ebiten.Image
}

type Enemies struct {
	game    *Game
	enemies []Enemy
}

func SpawnEnemies(game *Game) *Enemies {
	return &Enemies{
		game,
		firstEnemies,
	}
}

var firstEnemies = []Enemy{
	{
		40,
		50,
		75,
		assets.Spider,
	},
	{
		90,
		80,
		50,
		assets.GigaRat,
	},
}

func (_e *Enemies) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	for i := 0; i < len(firstEnemies); i++ {
		op.GeoM.Translate(float64(firstEnemies[i].X), float64(firstEnemies[i].Y))
		screen.DrawImage(firstEnemies[i].model, op)
	}
}
