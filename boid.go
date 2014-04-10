package boids

import (
	"fmt"
	"math/rand"
	//"math"
)

type BoidMap struct {
	Height int
	Width  int
}

//Holds state of a boid
type Boid struct {
	Location     PVector
	Velocity     PVector
	Acceleration PVector
	R            float64
	MaxForce     float64 // Maximum steering force
	MaxSpeed     float64 // Maximum speed
}

//Flock of boids
type Flock struct {
	Boids []Boid
}

//Creates a new Flock
func NewFlock() Flock {
	flock := Flock{}
	flock.Boids = make([]Boid, 10)
	for n := 0; n < 10; n++ {
		flock.Boids[n] = NewBoid(float64(rand.Intn(600)), float64(rand.Intn(400)))
	}
	return flock
}

//Run 1 step on flock
func (flock *Flock) Run(bMap BoidMap) {
	fmt.Println("starting...")
	for _, boid := range flock.Boids {
		//TODO: list should exclude current boid
		boid.Run(flock.Boids, bMap)
	}

	for h := 0; h < bMap.Height; h++ {
		for w := 0; w < bMap.Width; w++ {
			hit := false
			for _, boid := range flock.Boids {
				if int(boid.Location.X) == w && int(boid.Location.Y) == h {
					hit = true
				}
			}
			if hit {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("done")
}

//Creates a new Boid
func NewBoid(x float64, y float64) Boid {
	result := Boid{}
	result.Acceleration = NewPVector2D(0, 0)
	result.Velocity = NewRandom2dPVector()
	result.Location = NewPVector2D(x, y)
	result.R = 2.0
	result.MaxSpeed = 2
	result.MaxForce = 0.03
	return result
}

func RunSimulation() {

}

//Run 1 step in simulation
func (boid *Boid) Run(neighbours []Boid, bMap BoidMap) {
	boid.Flock(neighbours)
	boid.Update()
	boid.Wrap(bMap)
	//boid.Render()
}

//Render the boid
// func (boid *Boid) Render(bMap BoidMap) {
// 	for w := 0; w < bMap.Width; w++ {
// 		for h := 0; h < bMap.Height; h++ {

// 		}
// 	}
// }

//Wrap location when hitting edge of map
func (boid *Boid) Wrap(bMap BoidMap) {
	if boid.Location.X < -boid.R {
		boid.Location.X = float64(bMap.Width) + boid.R
	}
	if boid.Location.Y < -boid.R {
		boid.Location.Y = float64(bMap.Height) + boid.R
	}
	if boid.Location.X > float64(bMap.Width)+boid.R {
		boid.Location.X = -boid.R
	}
	if boid.Location.Y > float64(bMap.Height)+boid.R {
		boid.Location.Y = -boid.R
	}
}

//
func (boid *Boid) ApplyForce(force PVector) {
	boid.Acceleration.Add(force)
}

//Compute new acceleration value based on the 3
// rules (Separation, Alignment, Cohesion)
func (boid *Boid) Flock(neighbours []Boid) {
	sep := boid.Separate(neighbours)
	aln := boid.Align(neighbours)
	coh := boid.Cohesion(neighbours)
	// Weight forces
	sep.Mult(1.5)
	aln.Mult(1.0)
	coh.Mult(1.0)
	//Add forces to boids acceleration
	boid.ApplyForce(sep)
	boid.ApplyForce(aln)
	boid.ApplyForce(coh)
}

//Steers boid towards specified target
func (boid *Boid) Seek(target PVector) PVector {
	desired := target.Diff(boid.Location)
	desired.Normalize()
	desired.Mult(boid.MaxSpeed)

	steer := desired.Diff(boid.Velocity)
	steer.Limit(boid.MaxForce)
	return steer
}

//Calculates steering vector towards center of all neighbour boids
func (boid *Boid) Cohesion(neighbours []Boid) PVector {
	neighbourDist := 50.0
	sum := NewPVector2D(0, 0)
	count := 0

	for _, neighbour := range neighbours {
		d := boid.Location.Dist(neighbour.Location)
		if d > 0.0 && d < neighbourDist {
			sum.Add(neighbour.Location)
			count++
		}
	}

	if count > 0 {
		sum.Div(float64(count))
		return boid.Seek(sum)
	} else {
		return NewPVector2D(0, 0)
	}
}

//Aligns boid with neighbouring boids
func (boid *Boid) Align(neighbours []Boid) PVector {
	neighbourDist := 50.0
	sum := NewPVector2D(0, 0)
	count := 0

	for _, neighbour := range neighbours {
		d := boid.Location.Dist(neighbour.Location)
		if (d > 0.0) && d < neighbourDist {
			sum.Add(neighbour.Velocity)
			count++
		}
	}

	if count > 0 {
		sum.Div(float64(count))
		// Steering = Desired - Velocity
		sum.Normalize()
		sum.Mult(boid.MaxSpeed)
		steer := sum.Diff(boid.Velocity)
		steer.Limit(boid.MaxForce)
		return steer
	} else {
		return NewPVector2D(0, 0)
	}

}

//Steers boid away from neighbours to prevent collisions
// trys to maintain minSpace distance from neighbours
func (boid *Boid) Separate(boids []Boid) PVector {
	minSpace := 25.0
	steer := NewPVector2D(0, 0)
	count := 0

	for _, neighbour := range boids {
		d := boid.Location.Dist(neighbour.Location)
		// If the distance is greater than 0 (yourself)
		// and less than min desired distance
		if (d > 0) && (d < minSpace) {
			// Calculate vector pointing away from neighbour
			diff := boid.Location.Diff(neighbour.Location)
			diff.Normalize()
			diff.Div(d) // Weight by distance
			steer.Add(diff)
			count++
		}
	}

	// calc average of added vectors
	if count > 0 {
		steer.Div(float64(count))
	}

	// As long as the vector is greater than 0
	if steer.Mag() > 0 {
		//steering = desired - velocity
		steer.Normalize()
		steer.Mult(boid.MaxSpeed)
		steer = steer.Diff(boid.Velocity)
		steer.Limit(boid.MaxForce)
	}

	return steer
}

//Updates a boids location on map
func (boid *Boid) Update() {
	// Update velocity
	boid.Velocity.Add(boid.Acceleration)
	// Limit speed
	boid.Velocity.Limit(boid.MaxSpeed)
	boid.Location.Add(boid.Velocity)
	// Reset accelertion to 0 each cycle
	boid.Acceleration.Mult(0)
}
