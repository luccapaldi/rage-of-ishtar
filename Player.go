// Luc Capaldi
// 2021-12-29

package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

// Player struct
type Player struct {
	image      *ebiten.Image
	xImg, yImg float64
	xPos, yPos float64
	speed      float64
	direction  [2]int
}

// TODO: Update player orientation (rotation) -- maybe just sword.
// TODO: Implement velocity and/or acceleration so direction is non-orthogonal

// Player method to update position
func (player *Player) updatePosition() {
	// Reset player direction
	player.direction[0], player.direction[1] = 0, 0

	// Update player position based on keyboard input
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		player.yPos -= player.speed
		player.direction[1] += -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		player.yPos += player.speed
		player.direction[1] += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.xPos -= player.speed
		player.direction[0] += -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.xPos += player.speed
		player.direction[0] += 1
	}

	fmt.Printf("%+d, %+d\n", player.direction[0], player.direction[1])
	//fmt.Printf("%f, %f\n", player.xPos, player.yPos)

	// Enforce periodic boundaries
	updatePeriodic(&player.xPos, &player.yPos, &player.xImg, &player.yImg)

}
