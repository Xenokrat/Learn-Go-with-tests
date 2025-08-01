package clockface

import (
	"time"
	"math"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)
// A Point represents a two-dimentional Cartesian coordinates
type Point struct {
	X, Y float64	
}

// Second hand is the unit vector of the second hand of an analog clock
// at time `t` represented as a Point.
func SecondHand(tm time.Time) Point {
	p := secondHandPoint(tm)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	return p
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) +
	       (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) +
		   (math.Pi / (hoursInHalfClock / float64(t.Hour() % hoursInClock)))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}
