package tdl_test

import (
	"testing"

	"github.com/joematpal/go-tdl"
)

func TestGetMyAge(t *testing.T) {
	if Learner.Age == 0 {
		t.Error("Please answer how old you are in questions.go")
	}
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Integer - Age",
			want: Learner.Age,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tdl.GetMyAge(); got != tt.want {
				t.Errorf("GetMyAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMyHeight(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Integer - Height",
			want: Learner.Height,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tdl.GetMyHeight(); got != tt.want {
				t.Errorf("GetMyHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
