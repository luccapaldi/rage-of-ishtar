package main

import (
//	"fmt"
)

type Boid struct {
	pos       Vector2D
	vel       Vector2D
	acc       Vector2D
	steer     Vector2D
	neighbors []Boid
}

func (boid *Boid) Update(boids []Boid) {
	//boid.UpdateNeighbors(boids)
	//boid.Cohesion()
	//boid.Alignment()
	//boid.Separation()
	boid.UpdatePhysics()
	boid.EnforcePeriodicity()
}

func (boid *Boid) UpdateNeighbors(boids []Boid) {
	boid.neighbors = nil
	for i, _ := range boids {
		boid.IsNeighborsWith(boids[i])
	}
}

func (boid *Boid) IsNeighborsWith(candidateBoid Boid) {
	distanceBetween := boid.pos
	distanceBetween.Subtract(candidateBoid.pos)

	if (distanceBetween.Magnitude() < PERCEPTION_RADIUS) && (distanceBetween.Magnitude() > 0.0) {
		boid.neighbors = append(boid.neighbors, candidateBoid)
	}
}

//func (boid *Boid) Cohesion() {
//
//}

func (boid *Boid) Alignment() {
	boid.steer.Clear()

	if len(boid.neighbors) > 0 {
		for i, _ := range boid.neighbors {
			boid.steer.Add(boid.neighbors[i].vel)
		}
		boid.steer.DivideByScalar(float64(len(boid.neighbors)))
		boid.steer.Subtract(boid.vel)

		if boid.steer.Magnitude() > MAX_STEER {
			boid.steer.Limit(MAX_STEER)
		}
	}
}

//func (boid *Boid) Separation() {
//
//}

func (boid *Boid) UpdatePhysics() {
	//boid.acc.Add(boid.steer)
	boid.vel.Add(boid.acc)
	boid.pos.Add(boid.vel)
}

func (boid *Boid) EnforcePeriodicity() {
	currPosX := boid.pos.components[0]
	currPosY := boid.pos.components[1]

	if currPosX > (WINDOW_SIZE_X - BOID_SIZE_X) {
		boid.pos.components[0] = 0.0
	} else if currPosX < 0.0 {
		boid.pos.components[0] = WINDOW_SIZE_X - BOID_SIZE_X
	}

	if currPosY > (WINDOW_SIZE_Y - BOID_SIZE_Y) {
		boid.pos.components[1] = 0.0
	} else if currPosY < 0.0 {
		boid.pos.components[1] = WINDOW_SIZE_Y - BOID_SIZE_Y
	}
}
