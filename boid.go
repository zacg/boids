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

//Render the boid
func (boid *Boid) Render() {

}

//Updates a boids state
func (boid *Boid) Update() {
	// Update velocity
	boid.Velocity.Add(boid.Acceleration)
	// Limit speed
	boid.Velocity.Limit(boid.Maxspeed)
	boid.Location.Add(boid.Velocity)
	// Reset accelertion to 0 each cycle
	boid.Acceleration.Mult(0)
}

//Applies a force to boid
func (boid *Boid) ApplyForce(force PVector) {
	boid.Acceleration.Add(force)
}
