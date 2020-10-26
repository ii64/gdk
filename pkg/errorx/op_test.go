package errorx

import (
	"errors"
	"reflect"
	"testing"
)

func TestGetOps(t *testing.T) {
	tests := []struct {
		name  string
		input error
		want  []Op
	}{
		{
			name:  "No error",
			input: nil,
			want:  nil,
		},
		{
			name:  "Standard error",
			input: errors.New("standard-error"),
			want:  nil,
		},
		{
			name: "1 layer",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err:     nil,
			},
			want: []Op{
				"userService.FindUserByID",
			},
		},
		{
			name: "2 layer with standard error",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err:     errors.New("standard-error"),
			},
			want: []Op{
				"userService.FindUserByID",
			},
		},
		{
			name: "2 layer",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err: &Error{
					Code:    Gateway,
					Message: "Gateway server error.",
					Op:      "accountGateway.FindUserByID",
					Err:     nil,
				},
			},
			want: []Op{
				"userService.FindUserByID",
				"accountGateway.FindUserByID",
			},
		},
		{
			name: "3 layer",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err: &Error{
					Code:    Gateway,
					Message: "Gateway server error.",
					Op:      "accountGateway.FindUserByID",
					Err: &Error{
						Code:    Unknown,
						Message: "Unknown error.",
						Op:      "io.Write",
						Err:     nil,
					},
				},
			},
			want: []Op{
				"userService.FindUserByID",
				"accountGateway.FindUserByID",
				"io.Write",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetOps(tt.input)
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
			t.Log(got)
		})
	}
}
