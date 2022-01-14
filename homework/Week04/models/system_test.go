package models

import "testing"

func TestElevatorSystem_RequestElevatorToFloor(t *testing.T) {
	type fields struct {
		TotalFloorNum    int
		Elevator         Elevator
		FloorGates       []FloorGate
		ElevatorDestList []int
	}
	type args struct {
		elevator Elevator
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{name: "电梯系统测试", fields: fields{
			TotalFloorNum:    3,
			Elevator:         Elevator{CurrentFloor: 1},
			ElevatorDestList: []int{2, 3},
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sys := ElevatorSystem{
				TotalFloorNum:    tt.fields.TotalFloorNum,
				Elevator:         tt.fields.Elevator,
				FloorGates:       tt.fields.FloorGates,
				ElevatorDestList: tt.fields.ElevatorDestList,
			}
			sys.RequestElevatorToFloor(&tt.args.elevator)
		})
	}
}
