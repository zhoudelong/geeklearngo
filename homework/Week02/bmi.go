package main

import "fmt"

/*
BMI 计算体脂率
BMI = 体重（公斤） / (身高*身高)(米)
体脂率: （1.2 * BMI + 0.23*年龄 - 5.4 -10.8*性别(男1女0)） / 100
*/

//全局变量声明
var (
	weight      float64 //体重(公斤)
	height      float64 //身高(米)
	age         int     //年龄
	sex         string  //性别
	name        string
	sexRate     float64       //性别BMI转化值
	bmi         float64       //bmi值
	fatRate     float64       //体脂率
	avearageFat float64 = 0.0 //平均体脂
	peoplecount int     = 0   //录入人数计数器
	//isContinue  bool          //循环采集是否继续标志
)

func main() {
	fmt.Println("BMI is run......")
	for {
		getInfo()
		peoplecount++
		// bmi
		bmi = weight / (height * height)
		// 体脂率
		fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*sexRate) / 100
		// 计算,  输出, bmi, 姓名、性别、身高、体重、年龄信息, 建议
		fmt.Println(*(&name))
		Calculator()

		fmt.Println("--------------------------------")
		// 记录总人数, 平均体脂率
		fmt.Printf("共录入%d人记录，平均体脂为%2.3f \n",
			peoplecount, (avearageFat*float64(peoplecount-1)+fatRate)/float64(peoplecount))

		// 达到3人退出
		if peoplecount >= 3 {
			break
		} else {
			continue
		}
	}
}

//获取用户输入信息
func getInfo() {
	fmt.Printf("姓名: ")
	fmt.Scanln(&name)
	fmt.Printf("性别: ")
	fmt.Scanln(&sex)
	fmt.Printf("身高(米): ")
	fmt.Scanln(&height)
	fmt.Printf("体重(公斤): ")
	fmt.Scanln(&weight)
	fmt.Printf("年龄：")
	fmt.Scanln(&age)
	if sex == "男" {
		sexRate = 1
	} else {
		sexRate = 0
	}
}

//计算用户输入信息
func Calculator() {
	fmt.Println("--------------------------------------")
	if sex == "男" {
		if age >= 18 && age <= 39 {
			if fatRate <= 0.1 {
				fmt.Printf("BMI:%2.3f\n 体脂率：%2.3f\n太瘦了\n", bmi, fatRate)
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n标准，太棒了\n", bmi, fatRate)
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n偏重，不影响把妹，加油\n", bmi, fatRate)
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Printf("BMI:%2.3f\n 体脂率：%2.3f\n稍微有点胖了，需要控制\n", bmi, fatRate)
			} else {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n危险！危险！危险！要减肥了\n", bmi, fatRate)
			}
		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.11 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n太瘦了\n", bmi, fatRate)
			} else if fatRate > 0.11 && fatRate <= 0.17 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n标准，太棒了\n", bmi, fatRate)
			} else if fatRate > 0.17 && fatRate <= 0.22 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n偏重，不影响把妹，加油\n", bmi, fatRate)
			} else if fatRate > 0.22 && fatRate <= 0.27 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n稍微有点胖了，需要控制\n", bmi, fatRate)
			} else {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n危险！危险！危险！要减肥了\n", bmi, fatRate)
			}
		} else {
			if fatRate <= 0.13 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n太瘦了\n", bmi, fatRate)
			} else if fatRate > 0.13 && fatRate <= 0.19 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n标准，太棒了\n", bmi, fatRate)
			} else if fatRate > 0.19 && fatRate <= 0.24 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n偏重，不影响把妹，加油\n", bmi, fatRate)
			} else if fatRate > 0.24 && fatRate <= 0.29 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n稍微有点胖了，需要控制\n", bmi, fatRate)
			} else {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n危险！危险！危险！要减肥了\n", bmi, fatRate)
			}
		}
	} else {
		if age >= 18 && age <= 39 {
			if fatRate <= 0.2 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n太瘦了\n", bmi, fatRate)
			} else if fatRate > 0.2 && fatRate <= 0.27 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n标准，太棒了\n", bmi, fatRate)
			} else if fatRate > 0.27 && fatRate <= 0.34 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n偏重，不影响，加油\n", bmi, fatRate)
			} else if fatRate > 0.34 && fatRate <= 0.39 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n稍微有点胖了，需要控制\n", bmi, fatRate)
			} else {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n危险！危险！危险！要减肥了\n", bmi, fatRate)
			}
		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.21 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n太瘦了\n", bmi, fatRate)
			} else if fatRate > 0.21 && fatRate <= 0.28 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n标准，太棒了\n", bmi, fatRate)
			} else if fatRate > 0.28 && fatRate <= 0.35 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n偏重，不影响，加油\n", bmi, fatRate)
			} else if fatRate > 0.35 && fatRate <= 0.40 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n稍微有点胖了，需要控制\n", bmi, fatRate)
			} else {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n危险！危险！危险！要减肥了\n", bmi, fatRate)
			}
		} else {
			if fatRate <= 0.22 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n太瘦了\n", bmi, fatRate)
			} else if fatRate > 0.22 && fatRate <= 0.29 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n标准，太棒了\n", bmi, fatRate)
			} else if fatRate > 0.29 && fatRate <= 0.36 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n偏重，不影响，加油\n", bmi, fatRate)
			} else if fatRate > 0.36 && fatRate <= 0.41 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n稍微有点胖了，需要控制\n", bmi, fatRate)
			} else {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n危险！危险！危险！要减肥了\n", bmi, fatRate)
			}
		}
	}
}
