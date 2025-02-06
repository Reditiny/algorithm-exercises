package main

import "math"

/**
判断一个点是否在矩形内部
矩形四个点坐标为 (x1, y1), (x2, y2), (x3, y3), (x4, y4)
*/

func isInsideRectangle(x1, y1, x2, y2, x3, y3, x4, y4, x, y float64) bool {
	// 坐标变换，使得矩形平行于坐标轴
	e1 := math.Abs(x2 - x1)
	e2 := math.Abs(y2 - y1)
	e3 := math.Sqrt(e1*e1 + e2*e2)

	sin := e1 / e3
	cos := e2 / e3

	// 选择对角定点即可
	x1p := x1*cos + y1*sin
	y1p := y1*cos - x1*sin
	x3p := x3*cos + y3*sin
	y3p := y3*cos - x3*sin
	xp := x*cos + y*sin
	yp := y*cos - x*sin

	return x1p <= xp && xp <= x3p && y1p <= yp && yp <= y3p
}
