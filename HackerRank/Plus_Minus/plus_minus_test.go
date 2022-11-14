package Plus_Minus

import "testing"

func Test_plusMinus(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				arr: []int32{1, 1, 0, -1, -1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plusMinus(tt.args.arr)
		})
	}
}
