package tests_test

import (
	"Calendar/Controller"
	"testing"

	"github.com/labstack/echo/v4"
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
			Controller.GetAllStudent(tt.args.c)
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
			Controller.PostStudent(tt.args.c)
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
			Controller.GetStudentByID(tt.args.c)
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
			Controller.DeleteStudentByID(tt.args.c)
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
			if got := Controller.GetStudentNamebyId(tt.args.id); got != tt.want {
				t.Errorf("GetStudentNamebyId() = %v, want %v", got, tt.want)
			}
		})
	}
}
