package game

import (
	"DungeonGame/assets"
	"DungeonGame/params"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	game  *Game
	X     int
	Y     int
	model *ebiten.Image
	speed int
	size  int
}

func NewPlayer(game *Game) *Player {
	model := assets.Hero

	return &Player{
		game:  game,
		X:     0,
		Y:     0,
		model: model,
		speed: 1,
		size:  16,
	}
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += p.speed
	}

	// Prevent the player from moving out of the map boundaries
	if p.X < 0 {
		p.X = 0
	}
	if p.X > params.ScreenWidth-p.size {
		p.X = params.ScreenWidth - p.size
	}
	if p.Y < 0 {
		p.Y = 0
	}
	if p.Y > params.ScreenHeight-p.size {
		p.Y = params.ScreenHeight - p.size
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.X), float64(p.Y))

	screen.DrawImage(p.model, op)
}
