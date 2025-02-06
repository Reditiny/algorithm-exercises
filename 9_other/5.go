package main

/**
判断一个点是否在三角形内部
三角形三个点坐标为 (x1, y1), (x2, y2), (x3, y3)
*/

func isInsideTriangle(x1, y1, x2, y2, x3, y3, x, y float64) bool {

	// 向量积同号
	vectorProduct1 := (x2-x1)*(y-y1) - (x-x1)*(y2-y1)
	vectorProduct2 := (x3-x2)*(y-y2) - (x-x2)*(y3-y2)
	vectorProduct3 := (x1-x3)*(y-y3) - (x-x3)*(y1-y3)

	return (vectorProduct1 >= 0 && vectorProduct2 >= 0 && vectorProduct3 >= 0) || (vectorProduct1 <= 0 && vectorProduct2 <= 0 && vectorProduct3 <= 0)
}
