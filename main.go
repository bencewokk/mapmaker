package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
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
	data [36][64]int
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

func arrayToDeclaration(data [36][64]int) string {
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

	if ebiten.IsKeyPressed(ebiten.Key1) {
		selected = 1
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		selected = 2
	} else if ebiten.IsKeyPressed(ebiten.Key3) {
		selected = 3
	}

	intmx, intmy := ebiten.CursorPosition()
	cursor.float_x, cursor.float_y = float32(intmx), float32(intmy)
	x, y := ptid(cursor)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		currentmap.data[y][x] = selected
	}

	for i := 0; i < 36; i++ {

		for j := 0; j < 64; j++ {

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
				float32(j)*screendivisor,
				float32(i)*screendivisor,
				screendivisor*lcamera.zoom,
				screendivisor*lcamera.zoom,
				currenttilecolor,
				false,
			)
		}
	}

	write()

	fmt.Println()

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return outsideWidth, outsideHeight
}

func write() {
	// Define the file name and content
	filename := "example.txt"
	// Open the file (or create it if it doesn't exist) in write-only mode
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close() // Ensure file is closed after writing is done

	// Write content to the file
	_, err = file.WriteString(arrayToDeclaration(currentmap.data))
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
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
