package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestHandler_GetBooks(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		in *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				DB: tt.fields.DB,
			}
			h.GetBooks(tt.args.in)
		})
	}
}
