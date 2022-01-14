package models

type Passenger struct {
	Name         string
	Location     Floor
	IsInElevator bool
}

func (p Passenger) PushButton(system *ElevatorSystem) {
	// 按了电梯系统的按钮后，将目标楼层队列最后放入一个目标楼层
	system.ElevatorDestList = append(system.ElevatorDestList, p.Location)
}
