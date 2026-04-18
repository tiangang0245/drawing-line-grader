package main

import (
	"encoding/json"
	"fmt"
)

// StrokePoint 定义单个笔触点的结构
type StrokePoint struct {
	X float64 `json:"x"` // x坐标
	Y float64 `json:"y"` // y坐标
	T int64   `json:"t"` // 时间戳(ms)，通常用 int64 存储时间
	P *float64 `json:"p,omitempty"` // 压感值(0-1, 可选)，使用指针或 omitempty 来表示可选
}

// CanvasSize 定义画布尺寸结构
type CanvasSize struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// StrokeData 定义完整的线条数据结构
type StrokeData struct {
	Points     []StrokePoint `json:"points"`      // 线条点序列
	DeviceType string        `json:"deviceType"`  // 设备类型 (mouse/pen/touch)
	CanvasSize CanvasSize    `json:"canvasSize"`  // 画布尺寸
}

func main() {
	// --- 示例：如何使用 ---

	// 1. 创建一个带压感的点
	pressure := 0.85
	p1 := StrokePoint{X: 100.5, Y: 200.0, T: 1678880000, P: &pressure}

	// 2. 创建一个无压感的点 (P 为 nil)
	p2 := StrokePoint{X: 150.0, Y: 210.0, T: 1678880050}

	// 3. 组装完整数据
	data := StrokeData{
		Points:     []StrokePoint{p1, p2},
		DeviceType: "pen",
		CanvasSize: CanvasSize{Width: 800, Height: 600},
	}

	// 4. 序列化为 JSON (模拟传给前端或存储)
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}