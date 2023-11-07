package tests_test

import (
	"Calendar/Models"
	dataaccess "Calendar/data-access"
	"reflect"
	"testing"
)

func TestStartDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataaccess.StartDB()
		})
	}
}

func TestStudentByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []Models.Student
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dataaccess.StudentByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("StudentByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StudentByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStudentByIDFromDB(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    Models.Student
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dataaccess.GetStudentByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStudentByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStudentByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddStudent(t *testing.T) {
	type args struct {
		student Models.Student
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dataaccess.AddStudent(tt.args.student)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllStudents(t *testing.T) {
	tests := []struct {
		name string
		want []Models.Student
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dataaccess.GetAllStudents(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
