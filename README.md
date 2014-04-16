boids
=====

Boids simulation written in Golang

The Boids simulation/algorithm was created by Chris Renolds in 1986. It simulates a flock of birds by defining a simple set of 3 rules the birds/boids have to follow. The 3 rules are: separation (space between neighbours), alignment (move in the avg direction of neighbours), and cohesion (move to middle of local cluster of neighbours). 
[pseudocode](http://www.vergenet.net/~conrad/boids/pseudocode.html)

[background](http://en.wikipedia.org/wiki/Boids)

###Usage

	import (
		"github.com/zacg/boids"
		"fmt"
		)

	game := boids.NewGame()

	for i := 0; i < 250; i++ {
		//Clear the screen and print game map
		fmt.Print("\x0c", game.Run())
		fmt.Println("Time:", i)
		//Slow things down a bit		
		time.Sleep(time.Second / 10)
	}

