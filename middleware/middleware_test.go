package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/techschool/simplebank/token"
	"reflect"
	"testing"
)

func Test_authMiddleware(t *testing.T) {
	type args struct {
		tokenMaker token.Maker
	}
	tests := []struct {
		name string
		args args
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := authMiddleware(tt.args.tokenMaker); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
