package util

import "testing"

func TestGetAmountFromPercentage(t *testing.T) {
	type args struct {
		percentage float64
		from       float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"Get amount of percentage",
			args{percentage: 50.00, from: 1000},
			500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAmountFromPercentage(tt.args.percentage, tt.args.from); got != tt.want {
				t.Errorf("GetAmountFromPercentage() = %v, want %v", got, tt.want)
			}
		})
	}
}
