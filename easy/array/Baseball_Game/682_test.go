package Baseball_Game

import "testing"

func Test_calPoints(t *testing.T) {
	type args struct {
		ops []string
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
				ops: []string{"5", "2", "C", "D", "+"},
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calPoints(tt.args.ops); got != tt.want {
				t.Errorf("calPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
