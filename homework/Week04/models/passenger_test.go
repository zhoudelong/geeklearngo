package models

import "testing"

func TestPassenger_PushButton(t *testing.T) {
	type fields struct {
		Name         string
		Location     Floor
		IsInElevator bool
	}
	type args struct {
		system *ElevatorSystem
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{name: "test 1", fields: fields{
			Name:         "",
			Location:     1,
			IsInElevator: false,
		},
			args: args{system: &ElevatorSystem{
				TotalFloorNum:    3,
				Elevator:         Elevator{},
				FloorGates:       nil,
				ElevatorDestList: nil,
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Passenger{
				Name:         tt.fields.Name,
				Location:     tt.fields.Location,
				IsInElevator: tt.fields.IsInElevator,
			}
			p.PushButton(tt.args.system)
		})
	}
}
