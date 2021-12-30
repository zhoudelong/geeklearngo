package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

func main() {

	// 1: 计算体脂BMI
	mybmi, err := gobmi.BMI(73, 1.75)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("mybmi = ", mybmi)

	// 2: 本地扩展BMI、FatRate
	myFatRate := gobmi.FatRate(mybmi, 26, "男")
	fmt.Println("myFatRate = ", myFatRate)

	// 3: 自己测试 两数求和
	fmt.Println(gobmi.Add(3, 4))

}
