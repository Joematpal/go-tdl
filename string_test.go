package tdl_test

import (
	"testing"

	"github.com/joematpal/go-tdl"
)

func TestGetString(t *testing.T) {
	if Learner.FirstName == "" {
		t.Error("Please provide your first name in questions.go")
	}
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "String - First Name",
			want: Learner.FirstName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tdl.GetMyFirstName(); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMyLastName(t *testing.T) {
	if Learner.LastName == "" {
		t.Error("Please provide your last name in questions.go")
	}
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "String - Last Name",
			want: Learner.LastName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tdl.GetMyLastName(); got != tt.want {
				t.Errorf("GetMyLastName() = %v, want %v", got, tt.want)
			}
		})
	}
}
