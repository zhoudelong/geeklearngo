package main

import (
	"fmt"
)

// 判断两条直线是否平行
/*
两点确定一条直线（两点不能重合）
01.录入直线1在X,Y的值
02.录入直线2在X,Y的值
03.录入斜率K1,K2(条件两条直线平行，斜率相等)
04.录入两条直线在Y轴上的点b1,b2(b1不等于了2)
05录入数据都符合，那么这两条直线平行，否则重新录入
*/

// 定义全局变量
var (
	x1, y1         float64
	x2, y2         float64
	k1, b1, k2, b2 float64
)

func GetlineFunc() (string, float64, float64) {
	var k, b float64

	fmt.Println("请输入第一个点的坐标: x1, y1: ")
	fmt.Scanln(&x1)
	fmt.Scanln(&y1)
	fmt.Println("你输入的坐标为: (x1, y1): ", x1, y1)
	//x1, y1 = 3.0, 4.0
	p1 := [2]float64{x1, y1} // 装到数组中
	fmt.Println(p1)

	fmt.Println("请输入第二个点的坐标: x2, y2: ")
	fmt.Scanln(&x2)
	fmt.Scanln(&y2)
	fmt.Println("你输入的坐标为: (x2, y2): ", x2, y2)
	//x2, y2 = 8.0, 5.0
	p2 := [2]float64{x2, y2}
	fmt.Println(p2)

	if p1 == p2 {
		fmt.Println("输入点(x,y)重合, 请重新输入....")
		return "输入点(x,y)重合, 请重新输入....", 0, 0
	} else if x1 == x2 {
		fmt.Println("输入x的值重合, 请重新输入....")
		return "输入x的值重合, 请重新输入....", 0, 0
	}
	k = (y2 - y1) / (x2 - x1)
	b = y1 - k*x1
	return "yes", k, b

}

func main() {

	for {
		var err1, err2 string
		err1, k1, b1 = GetlineFunc()
		fmt.Println(err1, "k1: ", k1, "b1: ", b1)

		err2, k2, b2 = GetlineFunc()
		fmt.Println(err2, "k2: ", k2, "b2: ", b2)

		if k1 == k2 && b1 == b2 {
			fmt.Println("直线重合....")
			continue
		} else if k1 == k2 && b1 != b2 {
			fmt.Println("直线平行....")
			break
		} else if err1 != "yes" {
			continue
		}
	}
}
