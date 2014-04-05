package boids

import (
	"math"
	"math/rand"
)

type PVector struct {
	X float64
	Y float64
	Z float64
}

func NewPVector2D(x float64, y float64) PVector {
	return PVector{x, y, 0}
}

func NewPVectorFromAngle(angle float64) PVector {
	result := PVector{}

	result.X = math.Cos(angle)
	result.Y = math.Sin(angle)
	result.Z = 0
	return result
}

func NewRandom2dPVector() PVector {
	return NewPVectorFromAngle(rand.Float64() * math.Pi * 2)
}

//Creates a new random 3d PVector
func NewRandom3dPVector() PVector {
	var angle float64 = 0
	var vz float64 = 0
	var result = PVector{}

	angle = rand.Float64() * math.Pi * 2
	result.Z = rand.Float64()*2 - 1

	result.X = (math.Sqrt(1-vz*vz) * math.Cos(angle))
	result.Y = (math.Sqrt(1-vz*vz) * math.Sin(angle))

	return result
}

//Decrements vector by 1
func (pvec *PVector) Sub() {
	pvec.X--
	pvec.Y--
	pvec.Z--
}

//Increments vector by 1
func (pvec *PVector) Inc() {
	pvec.X++
	pvec.Y++
	pvec.Z++
}
