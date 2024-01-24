package windmousego

import (
	"math"
	"math/rand"
)

type MouseSettings struct {
	StartX     float64
	StartY     float64
	EndX       float64
	EndY       float64
	Gravity    float64
	Wind       float64
	MinWait    float64
	MaxWait    float64
	MaxStep    float64
	TargetArea float64
}

func (settings MouseSettings) GeneratePoints() [][]float64 {
	if settings.Gravity < 1 {
		settings.Gravity = 1
	}

	if settings.MaxStep == 0 {
		settings.MaxStep = 0.01
	}

	var oldX float64
	var oldY float64
	var velocityX float64
	var velocityY float64
	var randomDist float64
	var step float64
	var dist float64
	var currentWait float64

	windX := float64(rand.Intn(10))
	windY := float64(rand.Intn(10))

	newX := math.Round(settings.StartX)
	newY := math.Round(settings.StartY)

	waitDiff := settings.MaxWait - settings.MinWait

	sqrt2 := math.Sqrt(2.0)
	sqrt3 := math.Sqrt(3.0)
	sqrt5 := math.Sqrt(5.0)

	dist = math.Hypot(settings.EndX-settings.StartX, settings.EndY-settings.StartY)

	var points [][]float64

	for dist > 1.0 {
		settings.Wind = math.Min(settings.Wind, dist)

		if dist >= settings.TargetArea {
			w := math.Floor(rand.Float64()*settings.Wind*2 + 1)
			windX = windX/sqrt3 + (w-settings.Wind)/sqrt5
			windY = windY/sqrt3 + (w-settings.Wind)/sqrt5
		} else {
			windX = windX / sqrt2
			windY = windY / sqrt2
			if settings.MaxStep < 3 {
				settings.MaxStep = float64(rand.Intn(3)) + 3.0
			} else {
				settings.MaxStep = settings.MaxStep / sqrt5
			}
		}

		velocityX += windX
		velocityY += windY
		velocityX = velocityX + (settings.Gravity*(settings.EndX-settings.StartX))/dist
		velocityY = velocityY + (settings.Gravity*(settings.EndY-settings.StartY))/dist

		if math.Hypot(velocityX, velocityY) > settings.MaxStep {
			randomDist = settings.MaxStep/2.0 + math.Floor((rand.Float64()*settings.MaxStep)/2)
			velocityMag := math.Hypot(velocityX, velocityY)
			velocityX = (velocityX / velocityMag) * randomDist
			velocityY = (velocityY / velocityMag) * randomDist
		}

		oldX = math.Round(settings.StartX)
		oldY = math.Round(settings.StartY)
		settings.StartX += velocityX
		settings.StartY += velocityY

		dist = math.Hypot(settings.EndX-settings.StartX, settings.EndY-settings.StartY)

		newX = math.Round(settings.StartX)
		newY = math.Round(settings.StartY)

		step = math.Hypot(settings.StartX-oldX, settings.StartY-oldY)

		wait := math.Round(waitDiff*(step/settings.MaxStep) + settings.MinWait)

		currentWait += wait

		if oldX != newX || oldY != newY {
			points = append(points, []float64{newX, newY, currentWait})
		}
	}

	endX := math.Round(settings.EndX)
	endY := math.Round(settings.EndY)

	if endX != newX || endY != newY {
		points = append(points, []float64{newX, newY, currentWait})
	}

	return points
}
