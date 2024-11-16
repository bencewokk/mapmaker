package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func writeMapData() {
	filename := "example.txt"

	// Open file for writing
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Write map data
	for _, row := range currentmap.data {
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
		_, err = fmt.Fprintf(file, "\n")
		if err != nil {
			fmt.Println("Error writing to the file:", err)
			return
		}
	}

	// Separator for sprites
	_, err = fmt.Fprintf(file, "---SPRITES---\n")
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}

	// Write sprite data
	for _, sprite := range currentmap.sprites {
		_, err = fmt.Fprintf(file, "%d,%f,%f\n", sprite.typeOf, sprite.pos.float_x, sprite.pos.float_y)
		if err != nil {
			fmt.Println("Error writing sprite to the file:", err)
			return
		}
	}

	fmt.Println("File written successfully!")
}
func readMapData() {
	filename := "map.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	isReadingSprites := false
	currentmap.sprites = nil // Clear existing sprites

	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Look for the sprite section header
		if line == "---SPRITES---" {
			isReadingSprites = true
			continue
		}

		// Process sprite data
		if isReadingSprites {
			// Split sprite data by commas
			values := strings.Split(line, ",")
			if len(values) != 3 {
				fmt.Println("Invalid sprite data:", line)
				continue
			}

			// Parse the values for the sprite
			typeOf, err := strconv.Atoi(strings.TrimSpace(values[0]))
			if err != nil {
				fmt.Println("Error parsing sprite type:", err)
				continue
			}

			floatX, err := strconv.ParseFloat(strings.TrimSpace(values[1]), 32)
			if err != nil {
				fmt.Println("Error parsing sprite X position:", err)
				continue
			}

			floatY, err := strconv.ParseFloat(strings.TrimSpace(values[2]), 32)
			if err != nil {
				fmt.Println("Error parsing sprite Y position:", err)
				continue
			}

			// Create the sprite and add it to the map
			sprite := createSprite(createPos(float32(floatX), float32(floatY)), typeOf)
			currentmap.sprites = append(currentmap.sprites, sprite)
		} else {
			// Process map data
			values := strings.Split(line, ",")
			for x, value := range values {
				value = strings.TrimSpace(value)
				if value == "" {
					continue
				}

				intValue, err := strconv.Atoi(value)
				if err != nil {
					fmt.Println("Error parsing map value:", err)
					return
				}

				currentmap.data[y][x] = intValue
			}
			y++
		}
	}

	// Check for any errors encountered during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	fmt.Println("File read successfully!")
}
