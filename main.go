package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

type gamemap struct {
	// map data (2D array)
	//
	// 0 = not decided, 1 = mountains, 2 = plains, 3 = hills, 4 = forests
	data [100][100]int
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
		selected = 1
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		selected = 2
	} else if ebiten.IsKeyPressed(ebiten.Key3) {
		selected = 3
	} else if ebiten.IsKeyPressed(ebiten.Key4) {
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

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton1) {

	}

	for i := 0; i < 100; i++ {

		for j := 0; j < 100; j++ {

			switch currentmap.data[i][j] {
			case 2:
				currenttilecolor = mlightgreen
			case 3:
				currenttilecolor = mbrown
			case 1:
				currenttilecolor = mdarkgray
			case 4:
				currenttilecolor = mdarkgreen
			default:
				currenttilecolor = uitransparent

			}

			vector.DrawFilledRect(
				screen,
				(float32(j)*screendivisor+lcamera.pos.float_x*lcamera.zoom)-screenWidth/2,
				(float32(i)*screendivisor+lcamera.pos.float_y*lcamera.zoom)-screenHeight/2,
				screendivisor*lcamera.zoom,
				screendivisor*lcamera.zoom,
				currenttilecolor,
				false,
			)
		}
	}

	fps := ebiten.CurrentFPS()
	fpsText := fmt.Sprintf("FPS: %.2f", fps)
	ebitenutil.DebugPrint(screen, fpsText)

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		write()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return outsideWidth, outsideHeight
}

func write() {
	// Define the file name
	filename := "example.txt"

	// Open the file (or create it if it doesn't exist) in write-only mode
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close() // Ensure file is closed after writing is done

	// Iterate over each row of the map data (assuming `currentmap.data` is 100x100)
	for _, row := range currentmap.data {
		// Write each cell in the row separated by ", "
		for j, value := range row {
			if j > 0 {
				_, err = fmt.Fprintf(file, ", ")
				if err != nil {
					fmt.Println("Error writing to the file:", err)
					return
				}
			}
			_, err = fmt.Fprintf(file, "%d", value)
			if err != nil {
				fmt.Println("Error writing to the file:", err)
				return
			}
		}
		// End the line after each row of numbers
		_, err = fmt.Fprintf(file, "\n")
		if err != nil {
			fmt.Println("Error writing to the file:", err)
			return
		}
	}

	fmt.Println("File written successfully!")
}

func main() {
	gameinit()
	ebiten.SetWindowSize(1600, 900)
	ebiten.SetWindowTitle("map maker")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
