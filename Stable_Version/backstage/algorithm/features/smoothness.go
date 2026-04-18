package features

import (
	"math"
	"line-quality-evaluator/types"
)

func CalculateSmoothness(points []types.StrokePoint) float64 {
	if len(points) < 3 {
		return 1.0
	}

	totalDirectionChange := 0.0

	for i := 2; i < len(points); i++ {
		v1x := points[i-1].X - points[i-2].X
		v1y := points[i-1].Y - points[i-2].Y
		v2x := points[i].X - points[i-1].X
		v2y := points[i].Y - points[i-1].Y

		dot := v1x*v2x + v1y*v2y
		mag1 := math.Sqrt(v1x*v1x + v1y*v1y)
		mag2 := math.Sqrt(v2x*v2x + v2y*v2y)

		if mag1 > 0 && mag2 > 0 {
			cosTheta := dot / (mag1 * mag2)
			if cosTheta > 1 {
				cosTheta = 1
			} else if cosTheta < -1 {
				cosTheta = -1
			}
			angle := math.Acos(cosTheta)
			totalDirectionChange += angle
		}
	}

	smoothness := 1 - math.Min(1, totalDirectionChange/math.Pi)
	return smoothness
}
