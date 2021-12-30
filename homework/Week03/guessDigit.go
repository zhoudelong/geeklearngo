package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//guessDigitGameV1()
	guessDigitGameV2(0, 100)

}
func guessDigitGameV2(left, right int) {
	// 递归实现, 根据提示, 最快找到心中的数, 每次折中
	guessed := (left + right) / 2
	fmt.Println("我猜是: ", guessed)
	var getinput string
	fmt.Print("如果高了，输入1，如果低了，输入0；对了，输入9：") // 不带回车
	fmt.Scanln(&getinput)
	switch getinput {
	case "1":
		if left == right {
			fmt.Println("你是不是改变注意了 ? ")
			return
		}
		guessDigitGameV2(left, guessed-1)
	case "0":
		if left == right {
			fmt.Println("你是不是改变注意了 ? ")
			return
		}
		guessDigitGameV2(guessed+1, right)
	case "9":
		fmt.Println("你心里想的数字是: ", guessed)
	}

}

func guessDigitGameV1() {
	// loop
	var i string
	fmt.Println("请想出你心里的一个100以内的数, 并记住它: ")
	fmt.Scanln(&i)
	num, _ := strconv.Atoi(string(i)) // string to int
	//fmt.Println(num, reflect.TypeOf(num))
	fmt.Println("你心里的数字是: ", num)

	flag := true
	reader := bufio.NewReader(os.Stdin) //
	for flag {
		data, _, _ := reader.ReadLine()            //
		command, err := strconv.Atoi(string(data)) //
		if err != nil {
			fmt.Println("格式不对，请输入数字")
			continue
		}
		fmt.Println("你输入的数字:", command)
		if command == num {
			flag = false
			fmt.Println("恭喜你，答对了~")
		} else if command > num {
			fmt.Println("你输入的数字大于心中的数字，别灰心！再来一次~")
		} else if command < num {
			fmt.Println("你输入的数字小于心中的数字，别灰心！再来一次~")
		}
	}
}
