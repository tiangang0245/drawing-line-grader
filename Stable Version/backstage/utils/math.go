package utils

import "math"

func Clamp(value, min, max float64) float64 {
	return math.Max(min, math.Min(max, value))
}

func Lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

func Distance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

func Angle(x1, y1, x2, y2 float64) float64 {
	return math.Atan2(y2-y1, x2-x1)
}

func Mean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func StandardDeviation(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	avg := Mean(values)
	squareDiffs := 0.0
	for _, v := range values {
		squareDiffs += math.Pow(v-avg, 2)
	}
	return math.Sqrt(squareDiffs / float64(len(values)))
}
