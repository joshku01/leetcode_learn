package Binary_Gap

import "testing"

func Test_binaryGap(t *testing.T) {
	type args struct {
		n int
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
				n: 22,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryGap(tt.args.n); got != tt.want {
				t.Errorf("binaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
