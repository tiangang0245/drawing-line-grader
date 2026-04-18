package visualization

import "fmt"

func ScoreToColor(score float64) string {
	if score >= 0.8 {
		return "rgb(34, 197, 94)"
	} else if score >= 0.6 {
		return "rgb(234, 179, 8)"
	} else if score >= 0.4 {
		return "rgb(249, 115, 22)"
	} else {
		return "rgb(239, 68, 68)"
	}
}

func ScoresToColors(scores []float64) []string {
	colors := make([]string, len(scores))
	for i, score := range scores {
		colors[i] = ScoreToColor(score)
	}
	return colors
}

func ScoreToHexColor(score float64) string {
	if score >= 0.8 {
		return "#22c55e"
	} else if score >= 0.6 {
		return "#eab308"
	} else if score >= 0.4 {
		return "#f97316"
	} else {
		return "#ef4444"
	}
}

func FormatHeatmapJSON(scores []float64) string {
	colors := ScoresToColors(scores)
	result := "["
	for i, color := range colors {
		result += fmt.Sprintf(`"%s"`, color)
		if i < len(colors)-1 {
			result += ","
		}
	}
	result += "]"
	return result
}
