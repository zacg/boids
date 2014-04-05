package boids

import (
//"math/rand"
)

//Holds state of a boid
type Boid struct {
	Location     PVector
	Velocity     PVector
	Acceleration PVector
	R            float64
	Maxforce     float64 // Maximum steering force
	Maxspeed     float64 // Maximum speed
}

//Flock of boids
type Flock struct {
	Boids []Boid
}

//Creates a new Flock
func NewFlock() Flock {
	return Flock{}
}

//Run 1 step on flock
func (flock *Flock) Run() {
	for _, boid := range flock.Boids {
		boid.Run()
	}
}

//Creates a new Boid
func NewBoid(x float64, y float64) Boid {
	result := Boid{}
	result.Acceleration = NewPVector2D(0, 0)
	result.Velocity = NewRandom2dPVector()
	result.Location = NewPVector2D(x, y)
	result.R = 2.0
	result.Maxspeed = 2
	result.Maxforce = 0.03
	return result
}

//Run 1 step in simulation
func (boid *Boid) Run() {

}
