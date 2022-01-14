package test_cases

import (
	"Week04/models"
	"log"
)

func RunCase2() {
	eleSys := models.ElevatorSystem{
		TotalFloorNum: 5,
		Elevator: models.Elevator{
			Name:          "三菱电梯",
			RunningStatus: "停止",
		},
		ElevatorDestList: []int{},
	}

	// Case 2
	log.Println("Case 2 =====")
	log.Printf("获取楼层数：%v", eleSys.TotalFloorNum)
	eleSys.Elevator.CurrentFloor = 1
	log.Printf("电梯当前在 %v 层", eleSys.Elevator.CurrentFloor)
	passenger1 := models.Passenger{
		Name:         "小明",
		Location:     3,
		IsInElevator: false,
	}
	passenger1.PushButton(&eleSys)
	log.Printf("%v在 %v 层按了电梯", passenger1.Name, passenger1.Location)
	log.Printf("目标楼层队列：%v", eleSys.ElevatorDestList)
	eleSys.RequestElevatorToFloor(&eleSys.Elevator)
	log.Printf("电梯%v", eleSys.Elevator.RunningStatus)
}
