package types

type StrokePoint struct {
	X float64
	Y float64
	T int64
	P *float64
}

type StrokeData struct {
	Points    []StrokePoint
	DeviceType string
	CanvasSize CanvasSize
}

type CanvasSize struct {
	Width  float64
	Height float64
}

type DimensionScores struct {
	Smoothness       float64
	Stability        float64
	SpeedConsistency float64
	PressureControl  float64
}

type FinalScore struct {
	Score       int
	Breakdown   DimensionScores
	HeatmapData []float64
	Suggestions []string
}

type ScoringWeights struct {
	Smoothness       float64
	Stability        float64
	SpeedConsistency float64
	PressureControl  float64
}

var DefaultWeights = ScoringWeights{
	Smoothness:       0.4,
	Stability:        0.3,
	SpeedConsistency: 0.2,
	PressureControl:  0.1,
}
