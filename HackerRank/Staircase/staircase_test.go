package Staircase

import "testing"

func Test_staircase(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				n: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			staircase(tt.args.n)
		})
	}
}
