package models

import (
	"Week04/funcs"
	"log"
	"time"
)

type Button struct {
	Location Floor
	IsPushed bool
}

type FloorGate struct {
	Location Floor
	Button   Button
}

type ElevatorSystem struct {
	TotalFloorNum    int
	Elevator         Elevator
	FloorGates       []FloorGate
	ElevatorDestList []int
}

func (sys *ElevatorSystem) RequestElevatorToFloor(elevator *Elevator) {
	destFloor := sys.ElevatorDestList[0]
	// 获取当前电梯的目标楼层后，将其在目标楼层队列中移除掉
	sys.ElevatorDestList = funcs.DeleteItem(sys.ElevatorDestList, 0)
	elevator.goToDestFloor(destFloor)
}

// Run 为后台一直运行做准备，还未上线
func (sys ElevatorSystem) Run() {
	for {
		if len(sys.ElevatorDestList) > 0 {
			sys.RequestElevatorToFloor(&sys.Elevator)
		} else {
			log.Println("没有人请求电梯，电梯待命中")
		}
		time.Sleep(1 * time.Second)

	}
}
