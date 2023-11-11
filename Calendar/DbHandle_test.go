package Calendar_test

import (
	"Calendar"
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
			Calendar.StartDB()
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
		want    []Calendar.Student
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calendar.StudentByNameFromDB(tt.args.name)
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

func TestGetByIDFromDB(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    Calendar.Student
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calendar.GetStudentByIDFromDB(tt.args.id)
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
		student Calendar.Student
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
			got, err := Calendar.AddStudentFromDB(tt.args.student)
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
		want []Calendar.Student
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calendar.GetAllStudentsFromDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
