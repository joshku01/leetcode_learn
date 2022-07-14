package Maximum_Product_Difference_Between_Two_Pairs

import "testing"

func Test_maxProductDifference(t *testing.T) {
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
				nums: []int{5, 6, 2, 7, 4},
			},
			want: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductDifference(tt.args.nums); got != tt.want {
				t.Errorf("maxProductDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
