package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"time"
)

type Flock struct {
	image *ebiten.Image
	boids []Boid
}

func (flock *Flock) CreateBoids(NUMBER_OF_BOIDS int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < NUMBER_OF_BOIDS; i++ {
		randPos := Vector2D{[]float64{float64(rand.Intn(WINDOW_SIZE_X)),
			float64(rand.Intn(WINDOW_SIZE_Y))}}

		randVel := Vector2D{[]float64{float64(rand.Intn(2*WINDOW_SIZE_X) - WINDOW_SIZE_X),
			float64(rand.Intn(2*WINDOW_SIZE_Y) - WINDOW_SIZE_Y)}}

		fmt.Printf("%v\n", randVel)
		randVel.Limit(MAX_SPEED)

		flock.boids = append(flock.boids, Boid{pos: randPos,
			vel:       randVel,
			acc:       Vector2D{[]float64{0.0, 0.0}},
			steer:     Vector2D{[]float64{0.1, 0.1}},
			neighbors: []Boid{}})
	}
}

func (flock *Flock) Update() {
	for i, _ := range flock.boids {
		flock.boids[i].Update(flock.boids)
	}
}

func (flock *Flock) Draw(screen *ebiten.Image) {
	for i, _ := range flock.boids {
		boidOp := &ebiten.DrawImageOptions{}
		boidOp.GeoM.Translate(flock.boids[i].pos.components[0], flock.boids[i].pos.components[1])
		screen.DrawImage(flock.image, boidOp)
	}
}
