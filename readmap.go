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
	for _, row := range globalGameState.currentmap.data {
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
	for _, sprite := range globalGameState.currentmap.sprites {
		_, err = fmt.Fprintf(file, "%d,%f,%f\n", sprite.typeOf, sprite.pos.float_x, sprite.pos.float_y)
		if err != nil {
			fmt.Println("Error writing sprite to the file:", err)
			return
		}
	}

	_, err = fmt.Fprintf(file, "---NODES---\n")
	// Write nodes
	for _, n := range globalGameState.currentmap.nodes {
		_, err = fmt.Fprintf(file, "NODE,%d,%f,%f\n", n.id, n.pos.float_x, n.pos.float_y)
		if err != nil {
			fmt.Printf("error writing node to the file: %v", err)
			return
		}
	}

	_, err = fmt.Fprintf(file, "---PATHS---\n")
	// Write paths
	for _, p := range globalGameState.currentmap.paths {
		_, err = fmt.Fprintf(file, "PATH,%d,%d,%f\n", p.nodeA.id, p.nodeB.id, p.cost)
		if err != nil {
			fmt.Printf("error writing path to the file: %v", err)
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
	isReadingNodes := false
	isReadingPaths := false

	if isReadingPaths {

	}

	// Clear existing data
	globalGameState.currentmap.sprites = nil
	globalGameState.currentmap.nodes = nil
	globalGameState.currentmap.paths = nil

	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Look for section headers
		switch line {
		case "---SPRITES---":
			isReadingSprites = true
			isReadingNodes = false
			isReadingPaths = false
			continue
		case "---NODES---":
			isReadingSprites = false
			isReadingNodes = true
			isReadingPaths = false
			continue
		case "---PATHS---":
			isReadingSprites = false
			isReadingNodes = false
			isReadingPaths = true
			continue
		}

		// Process data based on current section
		if isReadingSprites {
			// Process sprite data
			values := strings.Split(line, ",")
			if len(values) != 3 {
				fmt.Println("Invalid sprite data:", line)
				continue
			}

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

			sprite := createSprite(createPos(float32(floatX), float32(floatY)), typeOf)
			globalGameState.currentmap.sprites = append(globalGameState.currentmap.sprites, sprite)

		} else if isReadingNodes {
			// Process node data
			values := strings.Split(line, ",")
			fmt.Println(values)
			if len(values) != 4 {
				fmt.Println("Invalid node data:", line)
				continue
			}

			if strings.TrimSpace(values[0]) != "NODE" {
				fmt.Println("Unexpected node data:", line)
				continue
			}

			id, err := strconv.Atoi(strings.TrimSpace(values[1]))
			if err != nil {
				fmt.Println("Error parsing node ID:", err)
				if idForNode < id {
					idForNode = id + 1
				}
				continue
			}

			floatX, err := strconv.ParseFloat(strings.TrimSpace(values[2]), 32)
			if err != nil {
				fmt.Println("Error parsing node X position:", err)
				continue
			}

			floatY, err := strconv.ParseFloat(strings.TrimSpace(values[3]), 32)
			if err != nil {
				fmt.Println("Error parsing node Y position:", err)
				continue
			}

			node := createNode(id, createPos(float32(floatX), float32(floatY)))
			globalGameState.currentmap.nodes = append(globalGameState.currentmap.nodes, node)

		} else if isReadingPaths {
			// Process path data
			values := strings.Split(line, ",")
			if len(values) != 4 {
				// fmt.Println("Invalid path data:", line)
				continue
			}

			if strings.TrimSpace(values[0]) != "PATH" {
				// fmt.Println("Unexpected path data:", line)
				continue
			}

			nodeAID, err := strconv.Atoi(strings.TrimSpace(values[1]))
			if err != nil {
				// fmt.Println("Error parsing path nodeA ID:", err)
				continue
			}

			nodeBID, err := strconv.Atoi(strings.TrimSpace(values[2]))
			if err != nil {
				// fmt.Println("Error parsing path nodeB ID:", err)
				continue
			}

			cost, err := strconv.ParseFloat(strings.TrimSpace(values[3]), 32)
			if err != nil {
				// fmt.Println("Error parsing path cost:", err)
				continue
			}

			path := createPath(findNodeByID(nodeAID), findNodeByID(nodeBID), float32(cost))
			globalGameState.currentmap.paths = append(globalGameState.currentmap.paths, path)

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

				globalGameState.currentmap.data[y][x] = intValue
			}
			y++
		}
	}

	// Check for any errors encountered during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	globalGameState.currentmap.height, globalGameState.currentmap.width = 100, 100

	fmt.Println("File read successfully!")
}
