package New_Year_Chaos

import "testing"

func Test_minimumBribes(t *testing.T) {
	type args struct {
		q []int32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				q: []int32{2, 5, 1, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minimumBribes(tt.args.q)
		})
	}
}
