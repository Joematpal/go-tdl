package tdl_test

import (
	"testing"

	"github.com/joematpal/go-tdl"
)

func TestGetBirthTimeAsFloat(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "Float - Birth Time",
			want: Learner.GetBirthMilitaryTimeFloat(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tdl.GetBirthTimeAsFloat(Learner.GetBirthHour(), Learner.GetBirthMin()); got != tt.want {
				t.Errorf("GetBirthTimeAsFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
