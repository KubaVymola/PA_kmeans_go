package main

import (
	"math"
	"math/rand"
)

type vec2d struct {
    x float64
    y float64
}

func randVec2d(maxVal int) vec2d {
    return vec2d{
        x: rand.Float64() * float64(maxVal),
        y: rand.Float64() * float64(maxVal),
    }
}

func getDistance(a, b vec2d) float64 {
	return math.Sqrt((math.Pow(a.x - b.x, 2) + math.Pow(a.y - b.y, 2)))
}

func (a vec2d) add(b vec2d) vec2d {
	return vec2d{a.x + b.x, a.y + b.y}
}

func (a vec2d) div(divisor float64) vec2d {
	return vec2d{a.x / divisor, a.y / divisor}
}