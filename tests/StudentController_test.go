package Calendar_test

import (
	"github.com/Wetherish/CalendarGO/Calendar"

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
		t.Run(tt.name, func(_ *testing.T) {
			Calendar.GetAllStudent(tt.args.c)
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
		t.Run(tt.name, func(_ *testing.T) {
			Calendar.PostStudent(tt.args.c)
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
		t.Run(tt.name, func(_ *testing.T) {
			Calendar.GetStudentByID(tt.args.c)
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
		t.Run(tt.name, func(_ *testing.T) {
			Calendar.DeleteStudent(tt.args.c)
		})
	}
}
