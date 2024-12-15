package game

import (
	"DungeonGame/assets"
	"DungeonGame/params"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	tileSize      = 16
	mapTilesWidth = 20
)

var (
	gameMap    = assets.GameMap
	tilesetImg = assets.Tileset
)

func DrawMap(screen *ebiten.Image) {
	for index, tile := range gameMap.Layers[0].Tiles {
		if !tile.Nil {
			spriteRect := tile.Tileset.GetTileRect(tile.ID)
			tileImage := tilesetImg.SubImage(spriteRect).(*ebiten.Image)

			x := (index % mapTilesWidth) * tileSize
			y := (index / mapTilesWidth) * tileSize

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(tileImage, op)
		}
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()), params.ScreenWidth-65, 1)
}
