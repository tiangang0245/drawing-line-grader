package scoring

import (
	"math"
	"line-quality-evaluator/types"
	"line-quality-evaluator/algorithm/features"
)

func CalculateWeightedScore(scores types.DimensionScores, weights types.ScoringWeights) float64 {
	weightedScore := scores.Smoothness*weights.Smoothness +
		scores.Stability*weights.Stability +
		scores.SpeedConsistency*weights.SpeedConsistency +
		scores.PressureControl*weights.PressureControl

	return weightedScore
}

func MapToFinalScore(weightedScore float64) int {
	return int(math.Round(weightedScore * 100))
}

func GenerateSuggestions(scores types.DimensionScores) []string {
	suggestions := []string{}

	if scores.Smoothness < 0.6 {
		suggestions = append(suggestions, "线条方向变化较大，建议练习长直线和曲线的连贯性")
	}

	if scores.Stability < 0.6 {
		suggestions = append(suggestions, "线条抖动较明显，建议加强控笔练习，保持手腕稳定")
	}

	if scores.SpeedConsistency < 0.6 {
		suggestions = append(suggestions, "运笔速度变化较大，建议保持匀速运笔")
	}

	if scores.PressureControl < 0.6 {
		suggestions = append(suggestions, "压感控制不够稳定，建议练习不同压力下的线条")
	}

	if len(suggestions) == 0 {
		suggestions = append(suggestions, "线条质量良好，继续保持练习")
	}

	return suggestions
}

func EvaluateStroke(strokeData types.StrokeData) types.FinalScore {
	points := features.PreprocessStrokeData(strokeData)

	if len(points) < 3 {
		return types.FinalScore{
			Score: 0,
			Breakdown: types.DimensionScores{
				Smoothness:       0,
				Stability:        0,
				SpeedConsistency: 0,
				PressureControl:  0,
			},
			HeatmapData: []float64{},
			Suggestions: []string{"线条过短，无法评估，请绘制更长的线条"},
		}
	}

	scores := features.CalculateDimensionScores(points)
	weightedScore := CalculateWeightedScore(scores, types.DefaultWeights)
	finalScore := MapToFinalScore(weightedScore)
	suggestions := GenerateSuggestions(scores)

	heatmapData := GenerateHeatmapData(points, scores)

	return types.FinalScore{
		Score:       finalScore,
		Breakdown:   scores,
		HeatmapData: heatmapData,
		Suggestions: suggestions,
	}
}

func GenerateHeatmapData(points []types.StrokePoint, scores types.DimensionScores) []float64 {
	heatmapData := make([]float64, len(points))

	for i := 0; i < len(points); i++ {
		start := i - 2
		if start < 0 {
			start = 0
		}
		end := i + 3
		if end > len(points) {
			end = len(points)
		}

		localPoints := points[start:end]
		localSmoothness := features.CalculateSmoothness(localPoints)
		localStability := features.CalculateStability(localPoints)

		localScore := localSmoothness*0.6 + localStability*0.4
		heatmapData[i] = localScore
	}

	return heatmapData
}
