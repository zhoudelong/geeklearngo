package models

import "testing"

func TestElevator_descend(t *testing.T) {
	type fields struct {
		Name          string
		CurrentFloor  Floor
		RunningStatus string
		Destination   Floor
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{name: "test 1", fields: fields{
			Name:          "e 1",
			CurrentFloor:  3,
			RunningStatus: "stopped",
			Destination:   2,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elevator{
				Name:          tt.fields.Name,
				CurrentFloor:  tt.fields.CurrentFloor,
				RunningStatus: tt.fields.RunningStatus,
				Destination:   tt.fields.Destination,
			}
			e.descend()
			if e.CurrentFloor != 2 {
				t.Fatalf("预期：2，得到：%v", e.CurrentFloor)
			}
		})
	}
}

func TestElevator_goToDestFloor(t *testing.T) {
	type fields struct {
		Name          string
		CurrentFloor  Floor
		RunningStatus string
		Destination   Floor
	}
	type args struct {
		dest Floor
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name:   "test 1",
			fields: fields{CurrentFloor: 5},
			args:   args{dest: 3}},
		{
			name:   "test 2",
			fields: fields{CurrentFloor: 1},
			args:   args{dest: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elevator{
				Name:          tt.fields.Name,
				CurrentFloor:  tt.fields.CurrentFloor,
				RunningStatus: tt.fields.RunningStatus,
				Destination:   tt.fields.Destination,
			}
			e.goToDestFloor(tt.args.dest)
			if e.CurrentFloor != 3 {
				t.Fatalf("预期：3，得到：%v", e.CurrentFloor)
			}
		})
	}
}

func TestElevator_rise(t *testing.T) {
	type fields struct {
		Name          string
		CurrentFloor  Floor
		RunningStatus string
		Destination   Floor
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{name: "test 1", fields: fields{
			Name:          "",
			CurrentFloor:  4,
			RunningStatus: "",
			Destination:   5,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elevator{
				Name:          tt.fields.Name,
				CurrentFloor:  tt.fields.CurrentFloor,
				RunningStatus: tt.fields.RunningStatus,
				Destination:   tt.fields.Destination,
			}
			e.rise()
			if e.CurrentFloor != 5 {
				t.Fatalf("预期：5，得到：%v", e.CurrentFloor)
			}
		})
	}
}
