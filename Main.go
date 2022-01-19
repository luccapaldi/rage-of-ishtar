package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	boidImage *ebiten.Image
	err       error
	flock     Flock
)

const (
	NAME                     = "Rage of Ishtar"
	BOID_IMAGE_PATH          = "/home/lcapaldi/go/src/rage-of-ishtar/assets/zombie.png"
	NUMBER_OF_BOIDS          = 100
	WINDOW_SIZE_X            = 736.0
	WINDOW_SIZE_Y            = 479.0
	BOID_SIZE_X, BOID_SIZE_Y = 24.0, 24.0
	MAX_SPEED                = 1.0
	MAX_STEER                = 0.00001
	PERCEPTION_RADIUS        = 10.0
)

func init() {
	boidImage, _, err = ebitenutil.NewImageFromFile("assets/zombie.png")
	if err != nil {
		log.Fatal(err)
	}

	flock = Flock{image: boidImage,
		boids: []Boid{}}
	flock.CreateBoids(NUMBER_OF_BOIDS)
}

type Game struct{}

func (g *Game) Update() error {
	flock.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	flock.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(WINDOW_SIZE_X), int(WINDOW_SIZE_Y)
}

func main() {
	ebiten.SetWindowSize(int(WINDOW_SIZE_X), int(WINDOW_SIZE_Y))
	ebiten.SetWindowTitle(NAME)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
