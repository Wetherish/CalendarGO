package tests_test

import (
	"Calendar/Controller"
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
			Controller.GetAlbums(tt.args.c)
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
			Controller.PostAlbums(tt.args.c)
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
			Controller.GetAlbumsByID(tt.args.c)
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
			Controller.DeleteByID(tt.args.c)
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
