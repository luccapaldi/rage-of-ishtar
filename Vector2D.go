package main

import (
	"math"
)

type Vector2D struct {
	components []float64
}

func (vec *Vector2D) Limit(maxMagnitude float64) {
	scalingFactor := maxMagnitude / vec.Magnitude()
	vec.MultiplyByScalar(scalingFactor)
}

func (vec *Vector2D) DivideByScalar(scalar float64) {
	vec.components[0] /= scalar
	vec.components[1] /= scalar
}

func (vec *Vector2D) MultiplyByScalar(scalar float64) {
	vec.components[0] *= scalar
	vec.components[1] *= scalar
}

func (vec *Vector2D) Magnitude() float64 {
	return math.Sqrt(math.Pow(vec.components[0], 2) + math.Pow(vec.components[1], 2))
}

func (vec *Vector2D) Add(inputVec Vector2D) {
	vec.components[0] += inputVec.components[0]
	vec.components[1] += inputVec.components[1]
}

func (vec *Vector2D) Subtract(inputVec Vector2D) {
	vec.components[0] -= inputVec.components[0]
	vec.components[1] -= inputVec.components[1]
}

func (vec *Vector2D) Clear() {
	vec.components = nil
}
