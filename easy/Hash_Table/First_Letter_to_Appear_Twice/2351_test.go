package First_Letter_to_Appear_Twice

import "testing"

func Test_repeatedCharacter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				s: "abcdd",
			},
			want: 'd',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedCharacter(tt.args.s); got != tt.want {
				t.Errorf("repeatedCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
