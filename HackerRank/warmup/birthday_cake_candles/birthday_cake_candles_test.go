package birthday_cake_candles

import "testing"

func Test_birthdayCakeCandles(t *testing.T) {
	type args struct {
		candles []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				candles: []int32{4, 4, 1, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := birthdayCakeCandles(tt.args.candles); got != tt.want {
				t.Errorf("birthdayCakeCandles() = %v, want %v", got, tt.want)
			}
		})
	}
}
