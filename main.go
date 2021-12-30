// Luc Capaldi
// 2021-12-29

package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// TODO: Create zombie class
// TODO: Create seeker

// Define global variables
var (
	screenX, screenY float64
	err              error
	background       *ebiten.Image
	player           Player
)

// Helper function to implement periodic boundaries
func updatePeriodic(xPos, yPos, xImg, yImg *float64) {
	if *xPos > screenX-*xImg {
		*xPos = 0.0
	} else if *xPos < 0.0 {
		*xPos = screenX - *xImg
	}

	if *yPos > screenY-*yImg {
		*yPos = 0.0
	} else if *yPos < 0.0 {
		*yPos = screenY - *yImg
	}
}

// init runs once before game loop begins
func init() {
	// Define window size
	screenX, screenY = 736.0, 479.0

	// Load background image
	background, _, err = ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}

	// Load player image
	player.image, _, err = ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	// Instantiate player
	player = Player{player.image,
		32,
		32,
		float64(screenX/2 - (player.xImg / 2)),
		float64(screenY/2 - (player.yImg / 2)),
		4,
		[2]int{0, 0}}
}

type Game struct{}

// Update game state
func (g *Game) Update() error {
	player.updatePosition()
	return nil
}

// Draw current game state
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.DrawImage(background, nil)

	// Draw player at current position
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(player.xPos, player.yPos)
	screen.DrawImage(player.image, playerOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(screenX), int(screenY)
}

func main() {
	ebiten.SetWindowSize(int(screenX), int(screenY))
	ebiten.SetWindowTitle("Rage of Ishtar")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
