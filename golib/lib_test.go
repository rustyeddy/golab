package golib

import (
	"errors"
	"testing"
)

// TestIsError ensures the associated function works correctly
func TestIsError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"noerror", args{nil}, false},
		{"error", args{errors.New("this is an error")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsError(tt.args.err); got != tt.want {
				t.Errorf("IsError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOK(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"OK", args{nil}, true},
		{"Not OK", args{errors.New("encountered a problem")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOK(tt.args.err); got != tt.want {
				t.Errorf("IsOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"No Error", args{nil}, true},
		{"Error", args{errors.New("this is an error")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotError(tt.args.err); got != tt.want {
				t.Errorf("IsNotError() = %v, want %v", got, tt.want)
			}
		})
	}
}
