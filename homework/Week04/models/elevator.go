package models

import (
	"log"
	"time"
)

type Elevator struct {
	Name          string
	CurrentFloor  Floor
	RunningStatus string
	Destination   Floor
	DoorStatus    string
}

func (e *Elevator) rise() {
	e.RunningStatus = "上行"
	e.logElevator()
	time.Sleep(1 * time.Second)
	e.CurrentFloor++
	e.RunningStatus = "停止"
}

func (e *Elevator) descend() {
	e.RunningStatus = "下行"
	e.logElevator()
	time.Sleep(1 * time.Second)
	e.CurrentFloor--
	e.RunningStatus = "停止"
}

func (e Elevator) logElevator() {
	log.Printf("电梯%v，当前楼层：%v", e.RunningStatus, e.CurrentFloor)
}

func (e *Elevator) goToDestFloor(dest Floor) {
	e.Destination = dest
	for e.Destination != e.CurrentFloor {
		switch {
		case e.Destination > e.CurrentFloor:
			e.rise()
		case e.Destination < e.CurrentFloor:
			e.descend()
		default:
			log.Println("目标楼层是当前楼层")
		}
	}
	log.Printf("电梯已到达目的地，当前楼层：%v", e.CurrentFloor)
}

// 下列代码还未上线，为详细的控制准备

func (e *Elevator) OpenDoor() {
	e.DoorStatus = "open"
	log.Println("电梯已开门")
}

func (e *Elevator) CloseDoor() {
	e.DoorStatus = "closed"
	log.Println("电梯已关门")
}
