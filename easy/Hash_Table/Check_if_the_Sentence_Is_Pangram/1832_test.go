package Check_if_the_Sentence_Is_Pangram

import "testing"

func Test_checkIfPangram(t *testing.T) {
	type args struct {
		sentence string
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
				sentence: "leetcode",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPangram(tt.args.sentence); got != tt.want {
				t.Errorf("checkIfPangram() = %v, want %v", got, tt.want)
			}
		})
	}
}
