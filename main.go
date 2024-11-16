package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// read more in gamestate
type camera struct {
	pos pos

	//used in rendering and collision checking
	zoom float32
}

var lcamera camera

// Screen sizes
var (
	width, height = ebiten.Monitor().Size()
	screenWidth   = float32(width)
	screenHeight  = float32(height)

	screendivisor    float32
	intscreendivisor int
)

var currentmap gamemap

type sprite struct {
	pos     pos
	texture *ebiten.Image
	typeOf  int
}

type gamemap struct {
	// map data (2D array)
	//
	// 0 = not decided, 1 = mountains, 2 = plains, 3 = hills, 4 = forests
	data    [100][100]int
	texture [100][100]*ebiten.Image

	sprites []sprite
}

func gameinit() {

	ebiten.SetWindowTitle("rpg")

	screendivisor = 900 / 45
	intscreendivisor = 900 / 45
	lcamera.pos = createPos(float32(width)/2, float32(height)/2)

	lcamera.zoom = 1

}

type Game struct{}

func (g *Game) Update() error {

	return nil
}

func arrayToDeclaration(data [1000][1000]int) string {
	result := "data := [][]int{\n"
	for _, row := range data {
		result += "\t{"
		for j, value := range row {
			result += strconv.Itoa(value)
			if j < len(row)-1 {
				result += ", "
			}
		}
		result += "},\n"
	}
	result += "}"
	return result
}

var cursor pos
var selected int

func (g *Game) Draw(screen *ebiten.Image) {

	parseTextureAndSprites()

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		lcamera.pos.float_x -= 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		lcamera.pos.float_x += 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		lcamera.pos.float_y += 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		lcamera.pos.float_y -= 10
	}

	if ebiten.IsKeyPressed(ebiten.Key1) {
		selected = 2
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		selected = 3
	} else if ebiten.IsKeyPressed(ebiten.Key0) {
		selected = 0
	}

	intmx, intmy := ebiten.CursorPosition()
	cursor.float_x, cursor.float_y = float32(intmx)-lcamera.pos.float_x+screenWidth/2, float32(intmy)-lcamera.pos.float_y+screenHeight/2
	x, y := ptid(cursor)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		if x >= 0 && y >= 0 {
			currentmap.data[y][x] = selected
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton2) {
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			// Define a radius for deleting sprites
			const deleteRadius = 50.0

			// Calculate the cursor position in world coordinates
			cursorWorldX := cursor.float_x + lcamera.pos.float_x - screenWidth/2
			cursorWorldY := cursor.float_y + lcamera.pos.float_y - screenHeight/2

			// Iterate through the sprites
			for i := 0; i < len(currentmap.sprites); i++ {
				dx := currentmap.sprites[i].pos.float_x - cursorWorldX
				dy := currentmap.sprites[i].pos.float_y - cursorWorldY
				distanceSquared := dx*dx + dy*dy

				// Nullify the sprite if within the delete radius
				if distanceSquared <= deleteRadius*deleteRadius {
					currentmap.sprites[i].texture = nil
				}
			}
		} else {
			// Add a new sprite at the cursor position
			currentmap.sprites = append(
				currentmap.sprites,
				createSprite(createPos(cursor.float_x-40, cursor.float_y-80), 0),
			)
		}
	}

	//TODO redo this comment and make this into a function
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if currentmap.texture[i][j] != nil {

				drawTile(screen, currentmap.texture[i][j], i, j)

			}

		}

	}

	for i := 0; i < len(currentmap.sprites); i++ {
		//fmt.Println(len(currentmap.sprites))
		sort.Slice(currentmap.sprites, func(i, j int) bool {
			return currentmap.sprites[i].pos.float_y < currentmap.sprites[j].pos.float_y
		})
		drawSprite(screen, currentmap.sprites[i].texture, currentmap.sprites[i].pos)
	}

	fps := ebiten.CurrentFPS()
	fpsText := fmt.Sprintf("FPS: %.2f", fps)
	ebitenutil.DebugPrint(screen, fpsText)

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		writeMapData()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return outsideWidth, outsideHeight
}

func main() {
	gameinit()
	readMapData()
	ebiten.SetWindowSize(1600, 900)
	ebiten.SetWindowTitle("map maker")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
