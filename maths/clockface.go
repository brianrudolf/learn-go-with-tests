package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(curTime time.Time) Point {
	second := curTime.Second()

	angel := float64(second) * 6.333
	// sin(angel) = opposite / hypo
	hypo := 150.00 / math.Sin(angel)

	return Point{150, 60}
}
