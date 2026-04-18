package features

import "line-quality-evaluator/types"

func CalculateDimensionScores(points []types.StrokePoint) types.DimensionScores {
	smoothness := CalculateSmoothness(points)
	stability := CalculateStability(points)
	speedConsistency := CalculateSpeedConsistency(points)
	pressureControl := CalculatePressureControl(points)

	return types.DimensionScores{
		Smoothness:       smoothness,
		Stability:        stability,
		SpeedConsistency: speedConsistency,
		PressureControl:  pressureControl,
	}
}
