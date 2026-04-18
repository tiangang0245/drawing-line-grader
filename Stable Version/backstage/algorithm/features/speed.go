package features

import (
	"math"
	"line-quality-evaluator/types"
)

func CalculateSpeedConsistency(points []types.StrokePoint) float64 {
	if len(points) < 2 {
		return 1.0
	}

	speeds := []float64{}

	for i := 1; i < len(points); i++ {
		dx := points[i].X - points[i-1].X
		dy := points[i].Y - points[i-1].Y
		dt := float64(points[i].T - points[i-1].T)

		if dt > 0 {
			distance := math.Sqrt(dx*dx + dy*dy)
			speed := distance / dt
			speeds = append(speeds, speed)
		}
	}

	if len(speeds) < 2 {
		return 1.0
	}

	sum := 0.0
	for _, s := range speeds {
		sum += s
	}
	meanSpeed := sum / float64(len(speeds))

	varianceSum := 0.0
	for _, s := range speeds {
		varianceSum += math.Pow(s-meanSpeed, 2)
	}
	speedStd := math.Sqrt(varianceSum / float64(len(speeds)))

	consistency := math.Max(0, 1-speedStd/(meanSpeed+0.001))
	return math.Min(1, consistency)
}

func CalculatePressureControl(points []types.StrokePoint) float64 {
	pressurePoints := []*types.StrokePoint{}
	for i := range points {
		if points[i].P != nil {
			pressurePoints = append(pressurePoints, &points[i])
		}
	}

	if len(pressurePoints) < 2 {
		return 1.0
	}

	totalPressureChange := 0.0
	for i := 1; i < len(pressurePoints); i++ {
		pressureDiff := math.Abs(*pressurePoints[i].P - *pressurePoints[i-1].P)
		totalPressureChange += pressureDiff
	}

	avgPressureChange := totalPressureChange / float64(len(pressurePoints)-1)
	pressureControl := math.Max(0, 1-avgPressureChange*5)
	return math.Min(1, pressureControl)
}
