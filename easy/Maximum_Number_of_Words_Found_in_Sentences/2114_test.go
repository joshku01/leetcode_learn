package Maximum_Number_of_Words_Found_in_Sentences

import "testing"

func Test_mostWordsFound(t *testing.T) {
	type args struct {
		sentences []string
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
				[]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostWordsFound(tt.args.sentences); got != tt.want {
				t.Errorf("mostWordsFound() = %v, want %v", got, tt.want)
			}
		})
	}
}
