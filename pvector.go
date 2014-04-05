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

//Calculates the length of the vector
func (pvec *PVector) Mag() float64 {
	return math.Sqrt(pvec.X*pvec.X + pvec.Y*pvec.Y + pvec.Z*pvec.Z)
}

//Calculates the squared magnitude of the vector
// (x*x + y*y + z*z)
func (pvec *PVector) MagSq() float64 {
	return pvec.X*pvec.X + pvec.Y*pvec.Y + pvec.Z*pvec.Z
}

//Limit the magnitude of vector to specified max
func (pvec *PVector) Limit(max float64) {
	if pvec.MagSq() > max*max {
		pvec.Normalize()
		pvec.Mult(max)
	}
}

//Adds 2 vectors
func (pvec *PVector) Add(pVector PVector) {
	pvec.X += pVector.X
	pvec.Y += pVector.Y
	pvec.Z += pVector.Z
}

//Divides the vector by the specified scalar
func (pvec *PVector) Div(n float64) {
	pvec.X /= n
	pvec.Y /= n
	pvec.Z /= n
}

//Multiplys vector by specified scalar
func (pvec *PVector) Mult(n float64) {
	pvec.X *= n
	pvec.Y *= n
	pvec.Z *= n
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

//Normalize vector to length 1
func (pvec *PVector) Normalize() {
	m := pvec.Mag()
	if m != 0 && m != 1 {
		pvec.Div(m)
	}
}
