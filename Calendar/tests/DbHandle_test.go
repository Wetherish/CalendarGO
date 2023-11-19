package Calendar_test

import (
	"Calendar"
	"reflect"
	"testing"
)

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
