package preprocessing

import (
	"math"
	"line-quality-evaluator/types"
)

func NormalizeCoordinates(points []types.StrokePoint, canvasSize types.CanvasSize) []types.StrokePoint {
	maxDimension := math.Max(canvasSize.Width, canvasSize.Height)
	normalized := make([]types.StrokePoint, len(points))
	for i, p := range points {
		normalized[i] = types.StrokePoint{
			X: p.X / maxDimension,
			Y: p.Y / maxDimension,
			T: p.T,
		}
		if p.P != nil {
			normalized[i].P = p.P
		}
	}
	return normalized
}

func FilterOutliers(points []types.StrokePoint, threshold float64) []types.StrokePoint {
	if len(points) < 3 {
		return points
	}

	distances := make([]float64, len(points)-1)
	for i := 1; i < len(points); i++ {
		dx := points[i].X - points[i-1].X
		dy := points[i].Y - points[i-1].Y
		distances[i-1] = math.Sqrt(dx*dx + dy*dy)
	}

	sum := 0.0
	for _, d := range distances {
		sum += d
	}
	mean := sum / float64(len(distances))

	varianceSum := 0.0
	for _, d := range distances {
		varianceSum += math.Pow(d-mean, 2)
	}
	std := math.Sqrt(varianceSum / float64(len(distances)))

	filtered := []types.StrokePoint{points[0]}
	for i := 1; i < len(points); i++ {
		dx := points[i].X - points[i-1].X
		dy := points[i].Y - points[i-1].Y
		distance := math.Sqrt(dx*dx + dy*dy)
		if distance <= mean+threshold*std {
			filtered = append(filtered, points[i])
		}
	}

	return filtered
}

func SmoothData(points []types.StrokePoint, windowSize int) []types.StrokePoint {
	if len(points) <= windowSize {
		return points
	}

	smoothed := make([]types.StrokePoint, len(points))
	halfWindow := windowSize / 2

	for i := 0; i < len(points); i++ {
		start := i - halfWindow
		if start < 0 {
			start = 0
		}
		end := i + halfWindow + 1
		if end > len(points) {
			end = len(points)
		}

		sumX, sumY, sumT := 0.0, 0.0, 0.0
		var sumP float64
		var countP int

		for j := start; j < end; j++ {
			sumX += points[j].X
			sumY += points[j].Y
			sumT += float64(points[j].T)
			if points[j].P != nil {
				sumP += *points[j].P
				countP++
			}
		}

		span := end - start
		smoothed[i] = types.StrokePoint{
			X: sumX / float64(span),
			Y: sumY / float64(span),
			T: int64(sumT / float64(span)),
		}
		if countP > 0 {
			smoothed[i].P = &sumP
			*smoothed[i].P /= float64(countP)
		}
	}

	return smoothed
}

func PreprocessStrokeData(strokeData types.StrokeData) []types.StrokePoint {
	points := strokeData.Points
	points = NormalizeCoordinates(points, strokeData.CanvasSize)
	points = FilterOutliers(points, 3.0)
	points = SmoothData(points, 3)
	return points
}
