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
	id         int
}

type path struct {
	nodeA node
	nodeB node
	cost  float32
}

func findNodeByID(id int) *node {
	for _, node := range globalGameState.currentmap.nodes {
		if node.id == id {
			return &node
		}
	}
	fmt.Printf("Warning: Node with ID %d not found\n", id)
	return nil
}
func createNode(id int, pos pos) node {
	return node{
		id:  id,
		pos: pos,
	}
}

func createPath(nodeA *node, nodeB *node, cost float32) path {
	if nodeA == nil || nodeB == nil {
		//	fmt.Println("Error: Cannot create path with nil nodes")
		return path{}
	}
	return path{
		nodeA: *nodeA,
		nodeB: *nodeB,
		cost:  cost,
	}
}

func drawPath(s *ebiten.Image, path path) {
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

var addigPath bool
var addigPathC2 bool
var idForNode int
var firstNode node

func createPathOnClick() {

	go func() {
		var posA, posB pos
		var nodeA, nodeB *node

		addigPath = true
		for {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) && ebiten.IsKeyPressed(ebiten.KeyControl) {
				posA = cursor
				break
			}
		}

		// Check the nearest node for posA
		nearestA, foundA := findNearestNode(posA, globalGameState.currentmap.nodes)
		if foundA {
			firstNode = *nearestA

			nodeA = nearestA
			nodeA.id = nearestA.id
		} else {
			nodeA = &node{pos: posA}
			nodeA.id = idForNode
			idForNode++
			globalGameState.currentmap.nodes = append(globalGameState.currentmap.nodes, *nodeA)

			firstNode = node{pos: posA}
		}

		addigPathC2 = true

		for {

			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) && ebiten.IsKeyPressed(ebiten.KeyAlt) {
				posB = cursor
				break
			}
		}

		// Check the nearest node for posB
		nearestB, foundB := findNearestNode(posB, globalGameState.currentmap.nodes)
		if foundB {
			nodeB = nearestB
			nodeB.id = nearestB.id
		} else {
			nodeB = &node{pos: posB}
			nodeB.id = idForNode + 1
			idForNode++
			globalGameState.currentmap.nodes = append(globalGameState.currentmap.nodes, *nodeB)

		}

		// Update availability of nodes
		if !nodeExistsInAvailables(nodeA, *nodeB) {
			nodeA.availables = append(nodeA.availables, *nodeB)
		}
		if !nodeExistsInAvailables(nodeB, *nodeA) {
			nodeB.availables = append(nodeB.availables, *nodeA)
		}

		// Create a path between the two nodes
		path := path{nodeA: *nodeA, nodeB: *nodeB, cost: Distance(nodeA.pos, nodeB.pos)}
		globalGameState.currentmap.paths = append(globalGameState.currentmap.paths, path)

		addigPath, addigPathC2 = false, false
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
