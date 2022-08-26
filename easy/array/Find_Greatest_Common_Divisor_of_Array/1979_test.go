package Find_Greatest_Common_Divisor_of_Array

import "testing"

func Test_findGCD(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{6, 7, 9},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGCD(tt.args.nums); got != tt.want {
				t.Errorf("findGCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
