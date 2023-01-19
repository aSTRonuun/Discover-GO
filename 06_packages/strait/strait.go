package main

import "math"

// Initializing with capital letter is public (visible in and out package)!

// Initializing with tiny letter is private (visible only in package)

// Point represent a coordernad in certisian plane
type Point struct {
	x float64
	y float64
}

func catets(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// The distance is responsible to calcule the linear distance  between two points
func Distance(a, b Point) float64 {
	cx, cy := catets(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
