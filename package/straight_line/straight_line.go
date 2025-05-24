package main

import "math"

// Starts with uppercase letter: Public
// Starts with lowercase letter: Private

type Point struct {
  x float64
  y float64
}

func legs(a, b Point) (lx, ly float64) {
  lx = math.Abs(a.x - b.x)
  ly = math.Abs(a.y - b.y)
  return
}

func Distance(a, b Point) float64 {
  lx, ly := legs(a, b)
  return math.Sqrt(math.Pow(lx, 2) + math.Pow(ly, 2))
}
