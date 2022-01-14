package test_cases

import (
	"Week04/models"
	"log"
)

func RunCase3() {
	eleSys := models.ElevatorSystem{
		TotalFloorNum: 5,
		Elevator: models.Elevator{
			Name:          "三菱电梯",
			RunningStatus: "停止",
		},
		ElevatorDestList: []int{},
	}

	// Case 3
	log.Println("Case 3.1 =====")
	log.Printf("获取楼层数：%v", eleSys.TotalFloorNum)
	eleSys.Elevator.CurrentFloor = 3
	log.Printf("电梯当前在 %v 层", eleSys.Elevator.CurrentFloor)
	eleSys.ElevatorDestList = []int{4, 2}
	for len(eleSys.ElevatorDestList) > 0 {
		log.Printf("目标楼层队列：%v", eleSys.ElevatorDestList)
		eleSys.RequestElevatorToFloor(&eleSys.Elevator)
	}
	log.Printf("电梯%v", eleSys.Elevator.RunningStatus)

	log.Println("Case 3.2 =====")
	log.Printf("获取楼层数：%v", eleSys.TotalFloorNum)
	eleSys.Elevator.CurrentFloor = 3
	log.Printf("电梯当前在 %v 层", eleSys.Elevator.CurrentFloor)
	eleSys.ElevatorDestList = []int{4, 5, 2}
	for len(eleSys.ElevatorDestList) > 0 {
		log.Printf("目标楼层队列：%v", eleSys.ElevatorDestList)
		eleSys.RequestElevatorToFloor(&eleSys.Elevator)
	}
	log.Printf("电梯%v", eleSys.Elevator.RunningStatus)
}
