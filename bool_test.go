package tdl_test

import (
	"testing"

	"github.com/joematpal/go-tdl"
)

func TestGetHasSiblings(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Bool - Has Siblings",
			want: Learner.Siblings,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tdl.GetHasSiblings(); got != tt.want {
				t.Errorf("GetHasSiblings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSiblingsFuncCheck(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Bool - Passed GetHasSiblings test",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tdl.SiblingsFuncCheck(); got != tt.want {
				t.Errorf("SiblingsFuncCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
