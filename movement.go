package main

//
import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var cursor pos

func checkMovementAndInput() {

	intmx, intmy := ebiten.CursorPosition()
	cursor.float_x, cursor.float_y =
		coffsetsx(float32(intmx)),
		coffsetsy(float32(intmy))

	x, y := ptid(cursor)

	// Handle movement based on key presses and check next tile for collisions
	if ebiten.IsKeyPressed(ebiten.KeyD) { // Move right
		char.pos.float_x += char.speed * float32(globalGameState.deltatime)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) { // Move left
		char.pos.float_x -= char.speed * float32(globalGameState.deltatime)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) { // Move up
		char.pos.float_y -= char.speed * float32(globalGameState.deltatime)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) { // Move down
		char.pos.float_y += char.speed * float32(globalGameState.deltatime)
	}

	if ebiten.IsKeyPressed(ebiten.Key1) {
		selected = 2
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		selected = 3
	} else if ebiten.IsKeyPressed(ebiten.Key0) {
		selected = 0
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		if x >= 0 && y >= 0 {
			globalGameState.currentmap.data[y][x] = selected
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		writeMapData()
	}

	// right click
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton2) && !ebiten.IsKeyPressed(ebiten.KeyShift) && !ebiten.IsKeyPressed(ebiten.KeyControl) {

		if globalGameState.pathcreatignmode {
			if !addigPath {
				go createPathOnClick()
			}

		} else {
			// Add a new sprite at the cursor position
			globalGameState.currentmap.sprites = append(
				globalGameState.currentmap.sprites,
				createSprite(createPos(cursor.float_x-40, cursor.float_y-80), 0))
		}
	}

	// right click + shift
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton2) && ebiten.IsKeyPressed(ebiten.KeyShift) && !ebiten.IsKeyPressed(ebiten.KeyControl) {
		if globalGameState.pathcreatignmode {
			globalGameState.pathcreatignmode = false
		} else {
			globalGameState.pathcreatignmode = true
		}
	}

}
