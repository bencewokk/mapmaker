package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Screen sizes
var (
	width, height = ebiten.Monitor().Size()
	screenWidth   = float32(width)
	screenHeight  = float32(height)

	screendivisor    float32
	intscreendivisor int
)

func gameinit() {

	load()
	readMapData()
	parseTextureAndSprites()

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("rpg")

	screendivisor = 30
	intscreendivisor = 30

	char.pos.float_y = screenHeight / 2
	char.pos.float_x = screenWidth / 2

}

type Game struct {
}

// Update method of the Game
func (g *Game) Update() error {
	checkMovementAndInput()
	return nil
}

var selected int

// Draw method of the Game
func (g *Game) Draw(screen *ebiten.Image) {

	now := time.Now()
	globalGameState.deltatime = now.Sub(globalGameState.lastUpdateTime).Seconds()
	globalGameState.lastUpdateTime = now

	updateCamera()

	parseTextureAndSprites()

	//TODO redo this comment and make this into a function
	for i := 0; i < globalGameState.currentmap.height; i++ {
		for j := 0; j < globalGameState.currentmap.width; j++ {
			if globalGameState.currentmap.texture[i][j] != nil {

				drawTile(screen, globalGameState.currentmap.texture[i][j], i, j)

			}

		}

	}

	for i := 0; i < len(globalGameState.currentmap.sprites); i++ {
		sort.Slice(globalGameState.currentmap.sprites, func(i, j int) bool {
			return globalGameState.currentmap.sprites[i].pos.float_y < globalGameState.currentmap.sprites[j].pos.float_y
		})
		drawSprite(screen, globalGameState.currentmap.sprites[i].texture, globalGameState.currentmap.sprites[i].pos)

	}

	for i := 0; i < len(globalGameState.currentmap.paths); i++ {
		drawPath(screen, globalGameState.currentmap.paths[i])
	}

	fmt.Println(globalGameState.currentmap.paths)
	// var op *ebiten.DrawImageOptions
	// screen.DrawImage(cutCam(screen, createPos(30, 30)), op)

	fps := ebiten.CurrentFPS()
	fpsText := fmt.Sprintf("FPS: %.2f", fps)
	ebitenutil.DebugPrint(screen, fpsText)
}

// Layout method of the Game
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	gameinit()
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
