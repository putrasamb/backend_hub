package util

import "testing"

func TestGetFloatingDecimal(t *testing.T) {
	type args struct {
		of        float64
		precision float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"Test zero value",
			args{of: (45) - 45, precision: 3},
			0.000,
		},

		{
			"Test floating number value",
			args{of: 30.1329 / 2, precision: 5},
			15.06645,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFloatingDecimal(tt.args.of, tt.args.precision); got != tt.want {
				t.Errorf("GetFloatingDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
