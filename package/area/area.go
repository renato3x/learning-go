package area

import "math"

const PI = 3.1415

func Circle(radius float64) float64 {
	return PI * math.Pow(radius, 2)
}

func Rectangle(length, width float64) float64 {
	return length * width
}

func _EquilateralTriangle(base, height float64) float64 {
	return base * height / 2
}
