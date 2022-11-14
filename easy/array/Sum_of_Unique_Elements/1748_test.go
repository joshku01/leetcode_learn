package Sum_of_Unique_Elements

import "testing"

func Test_sumOfUnique(t *testing.T) {
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
				nums: []int{1, 2, 3, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfUnique(tt.args.nums); got != tt.want {
				t.Errorf("sumOfUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
