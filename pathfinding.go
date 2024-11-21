package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type node struct {
	pos        pos
	availables []node
}

type path struct {
	nodeA node
	nodeB node
	cost  float32
}

func createPath(pointA, pointB node) path {
	return path{nodeA: pointA, nodeB: pointB, cost: Distance(pointA.pos, pointB.pos)}
}

func drawPath(s *ebiten.Image, path path) {
	fmt.Println("done4")
	ebitenutil.DrawLine(s, float64(offsetsx(path.nodeA.pos.float_x)), float64(offsetsy(path.nodeA.pos.float_y)),
		float64(offsetsx(path.nodeB.pos.float_x)), float64(offsetsy(path.nodeB.pos.float_y)), uidarkred)
}

const maxNodeDistance float32 = 10.0 // Maximum distance to snap to an existing node

func findNearestNode(pos pos, nodes []node) (*node, bool) {
	var nearest *node
	minDistance := maxNodeDistance
	for i := range nodes {
		dist := Distance(pos, nodes[i].pos)
		if dist < minDistance {
			minDistance = dist
			nearest = &nodes[i]
		}
	}
	return nearest, nearest != nil
}

func createPathOnClick() {
	go func() {
		var posA, posB pos
		var nodeA, nodeB *node

		fmt.Println("done1")
		for {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) && ebiten.IsKeyPressed(ebiten.KeyControlLeft) {
				posA = cursor
				break
			}
		}

		// Check the nearest node for posA
		nearestA, foundA := findNearestNode(posA, globalGameState.currentmap.nodes)
		if foundA {
			nodeA = nearestA
		} else {
			nodeA = &node{pos: posA}
			globalGameState.currentmap.nodes = append(globalGameState.currentmap.nodes, *nodeA)
		}

		fmt.Println("done2")
		for {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) && ebiten.IsKeyPressed(ebiten.KeyControlRight) {
				posB = cursor
				break
			}
		}

		// Check the nearest node for posB
		nearestB, foundB := findNearestNode(posB, globalGameState.currentmap.nodes)
		if foundB {
			nodeB = nearestB
		} else {
			nodeB = &node{pos: posB}
			globalGameState.currentmap.nodes = append(globalGameState.currentmap.nodes, *nodeB)
		}

		// Create a path between the two nodes
		fmt.Println("done3")
		path := path{nodeA: *nodeA, nodeB: *nodeB, cost: Distance(nodeA.pos, nodeB.pos)}
		globalGameState.currentmap.paths = append(globalGameState.currentmap.paths, path)

		// Update availability of nodes
		if !nodeExistsInAvailables(nodeA, *nodeB) {
			nodeA.availables = append(nodeA.availables, *nodeB)
		}
		if !nodeExistsInAvailables(nodeB, *nodeA) {
			nodeB.availables = append(nodeB.availables, *nodeA)
		}
	}()
}

// Check if a node exists in the availables list
func nodeExistsInAvailables(n *node, target node) bool {
	for _, available := range n.availables {
		if available.pos == target.pos {
			return true
		}
	}
	return false
}
