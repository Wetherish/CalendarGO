package Calendar

import (
	"testing"

	"github.com/labstack/echo/v5"
)

func TestGetStudent(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAllStudent(tt.args.c)
		})
	}
}

func TestPostStudent(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostStudent(tt.args.c)
		})
	}
}

func TestGetStudentByID(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetStudentByID(tt.args.c)
		})
	}
}

func TestDeleteStudentByID(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteStudentByID(tt.args.c)
		})
	}
}

func TestGetStudentNamebyId(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStudentNamebyId(tt.args.id); got != tt.want {
				t.Errorf("GetStudentNamebyId() = %v, want %v", got, tt.want)
			}
		})
	}
}
