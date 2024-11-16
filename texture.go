package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func parseTextureAndSprites() {

	var (
		DDDDD *ebiten.Image = loadPNG("import/tiles/Dry2Grass_DDDD.png")
		DDDDG *ebiten.Image = loadPNG("import/tiles/Dry2Grass_DDDG.png")
		DDDGD *ebiten.Image = loadPNG("import/tiles/Dry2Grass_DDGD.png")
		DDDGG *ebiten.Image = loadPNG("import/tiles/Dry2Grass_DDGG.png")
		DDGDD *ebiten.Image = loadPNG("import/tiles/Dry2Grass_DGDD.png")
		DDGDG *ebiten.Image = loadPNG("import/tiles/Dry2Grass_DGDG.png")
		DDGGD *ebiten.Image = loadPNG("import/tiles/Dry2Grass_DGGD.png")
		DDGGG *ebiten.Image = loadPNG("import/tiles/Dry2Grass_DGGG.png")
		DGDDD *ebiten.Image = loadPNG("import/tiles/Dry2Grass_GDDD.png")
		DGDDG *ebiten.Image = loadPNG("import/tiles/Dry2Grass_GDDG.png")
		DGDGD *ebiten.Image = loadPNG("import/tiles/Dry2Grass_GDGD.png")
		DGDGG *ebiten.Image = loadPNG("import/tiles/Dry2Grass_GDGG.png")
		DGGDD *ebiten.Image = loadPNG("import/tiles/Dry2Grass_GGDD.png")
		DGGDG *ebiten.Image = loadPNG("import/tiles/Dry2Grass_GGDG.png")
		DGGGD *ebiten.Image = loadPNG("import/tiles/Dry2Grass_GGGD.png")
		DGGGG *ebiten.Image = loadPNG("import/tiles/Dry2Grass_GGGG.png")
	)

	var dryTransitionsTextures = map[string]*ebiten.Image{
		"DDDD": DDDDD,
		"DDDG": DDDDG,
		"DDGD": DDDGD,
		"DDGG": DDDGG,
		"DGDD": DDGDD,
		"DGDG": DDGDG,
		"DGGD": DDGGD,
		"DGGG": DDGGG,
		"GDDD": DGDDD,
		"GDDG": DGDDG,
		"GDGD": DGDGD,
		"GDGG": DGDGG,
		"GGDD": DGGDD,
		"GGDG": DGGDG,
		"GGGD": DGGGD,
		"GGGG": DGGGG,
	}

	var (
		Grass_S1 *ebiten.Image = loadPNG("import/tiles/Grass_S1.png")
		Grass_S2 *ebiten.Image = loadPNG("import/tiles/Grass_S2.png")
		Grass_S3 *ebiten.Image = loadPNG("import/tiles/Grass_S3.png")
		Grass_S6 *ebiten.Image = loadPNG("import/tiles/Grass_S6.png")
		Grass_S8 *ebiten.Image = loadPNG("import/tiles/Grass_S8.png")

		Grass_S4 *ebiten.Image = loadPNG("import/tiles/Grass_S4.png") // normal
		Grass_S5 *ebiten.Image = loadPNG("import/tiles/Grass_S5.png") // normal
		Grass_S7 *ebiten.Image = loadPNG("import/tiles/Grass_S7.png") // normal
	)

	var grassTextures = []*ebiten.Image{
		Grass_S1, Grass_S2, Grass_S3, Grass_S6, Grass_S8, Grass_S4, Grass_S5, Grass_S7,
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {

			var textureID string = ""

			if currentmap.data[i][j] == 3 {
				// Check for out-of-bounds for each neighboring tile
				if currentmap.data[i-1][j] == 3 { // upper
					textureID += "D"
				} else {
					textureID += "G"
				}

				if currentmap.data[i][j-1] == 3 { // left
					textureID += "D"
				} else {
					textureID += "G"
				}

				if currentmap.data[i][j+1] == 3 { // right
					textureID += "D"
				} else {
					textureID += "G"
				}

				if currentmap.data[i+1][j] == 3 { // lower
					textureID += "D"
				} else {
					textureID += "G"
				}

				// Assign the appropriate texture based on the ID
				if texture, exists := dryTransitionsTextures[textureID]; exists {
					currentmap.texture[i][j] = texture
				}
			} else if currentmap.data[i][j] == 2 {
				if calcChance(10) {
					currentmap.texture[i][j] = grassTextures[rand.Int31n(5)]
				} else {
					currentmap.texture[i][j] = grassTextures[rand.Int31n(3)+5]
				}
			}
		}
	}

}

var (
	tree_S1 *ebiten.Image = loadPNG("import/prop/tree1.png")
	tree_S2 *ebiten.Image = loadPNG("import/prop/tree2.png")
)

var trees = []*ebiten.Image{tree_S1, tree_S2}

func createSprite(pos pos, typeOf int) sprite {
	var texture *ebiten.Image
	rand.Seed(1)
	if typeOf == 0 {
		texture = trees[rand.Intn(len(trees))]
	}
	return sprite{pos: pos, texture: texture}
}

func drawSprite(screen, t *ebiten.Image, pos pos) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(1.5, 1.5)

	op.GeoM.Translate(
		float64(((pos.float_x)+lcamera.pos.float_x)-screenWidth/2),
		float64(((pos.float_y)+lcamera.pos.float_y)-screenHeight/2))
	screen.DrawImage(t, op)

}

func drawTile(screen, t *ebiten.Image, i, j int) {
	op := &ebiten.DrawImageOptions{}

	originalWidth, originalHeight := t.Size()
	scaleX := float64(screendivisor) / float64(originalWidth)
	scaleY := float64(screendivisor) / float64(originalHeight)
	op.GeoM.Scale(scaleX, scaleY)

	op.GeoM.Translate(
		float64((float32(j)*screendivisor+lcamera.pos.float_x*lcamera.zoom)-screenWidth/2),
		float64((float32(i)*screendivisor+lcamera.pos.float_y*lcamera.zoom)-screenHeight/2))
	screen.DrawImage(t, op)
}

func loadSrite(screen, s *ebiten.Image, p pos) {

}
