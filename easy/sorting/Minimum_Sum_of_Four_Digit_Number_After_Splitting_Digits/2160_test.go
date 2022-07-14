package Minimum_Sum_of_Four_Digit_Number_After_Splitting_Digits

import "testing"

func Test_minimumSum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				num: 2932,
			},
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSum(tt.args.num); got != tt.want {
				t.Errorf("minimumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
