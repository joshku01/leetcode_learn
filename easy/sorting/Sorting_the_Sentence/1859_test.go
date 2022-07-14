package Sorting_the_Sentence

import "testing"

func Test_sortSentence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				s: "is2 sentence4 This1 a3",
			},
			want: "This is a sentence",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortSentence(tt.args.s); got != tt.want {
				t.Errorf("sortSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
