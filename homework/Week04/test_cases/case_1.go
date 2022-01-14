package test_cases

import (
	"Week04/models"
	"log"
)

func RunCase1() {
	// 实例化一个电梯系统，楼层有 5 层
	eleSys := models.ElevatorSystem{
		TotalFloorNum: 5,
		Elevator: models.Elevator{
			Name:          "三菱电梯",
			RunningStatus: "停止",
		},
		ElevatorDestList: []int{},
	}

	// Case 1
	log.Println("Case 1 =====")
	log.Printf("获取楼层数：%v", eleSys.TotalFloorNum)
	log.Printf("目标楼层队列：%v", eleSys.ElevatorDestList)
	log.Printf("目前没有人请求电梯，电梯状态：%v", eleSys.Elevator.RunningStatus)
}
