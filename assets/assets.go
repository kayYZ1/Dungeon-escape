package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

//go:embed *
var assets embed.FS

var (
	Hero    = mustLoadImage("Tiles/tile_0098.png")
	Tileset = mustLoadImage("Tilemap/tilemap.png")
	GameMap = loadTmx("assets/Tiled/dungeonescape.tmx")
	Spider  = mustLoadImage("Tiles/tile_0122.png")
	GigaRat = mustLoadImage("Tiles/tile_0123.png")
)

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func loadTmx(name string) *tiled.Map {
	gMap, err := tiled.LoadFile(name)
	if err != nil {
		panic(err)
	}

	return gMap
}
