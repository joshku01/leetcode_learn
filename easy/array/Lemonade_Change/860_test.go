package Lemonade_Change

import "testing"

func Test_lemonadeChange(t *testing.T) {
	type args struct {
		bills []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				bills: []int{5, 5, 5, 10, 5, 20, 5, 10, 5, 20},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lemonadeChange(tt.args.bills); got != tt.want {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
