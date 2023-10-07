package Controller

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAlbums(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAlbums(tt.args.c)
		})
	}
}

func TestPostAlbums(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostAlbums(tt.args.c)
		})
	}
}

func TestGetAlbumsByID(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAlbumsByID(tt.args.c)
		})
	}
}

func TestDeleteByID(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteByID(tt.args.c)
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
