package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("计算圆的面积....")
	
	var r float64
	fmt.Printf("请输入园半径: ")

	for {
		_, err := fmt.Scanln(&r)
		if err != nil {
			fmt.Println("输入必须是数字类型, 请重新输入! ")
			continue
		}

		if r <= 0 {
			fmt.Println("圆半径必须大于等于1, 请重新输入! ")
			continue
		} else {
			fmt.Println("圆面积是: ", math.Pi*r*r)
			break
		}

	}
}
