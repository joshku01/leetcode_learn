package MIni_Max_Sum

import "testing"

func Test_miniMaxSum(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{
				arr: []int32{1, 8, 3, 5, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			miniMaxSum(tt.args.arr)
		})
	}
}
