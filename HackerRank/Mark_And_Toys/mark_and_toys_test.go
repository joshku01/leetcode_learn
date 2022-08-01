package Mark_And_Toys

import (
	"reflect"
	"testing"
)

func Test_maximumToys(t *testing.T) {
	type args struct {
		prices []int32
		k      int32
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
				prices: []int32{1, 12, 5, 111, 200, 1000, 10},
				k:      int32(50),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumToys(tt.args.prices, tt.args.k); got != tt.want {
				t.Errorf("maximumToys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeConversion(t *testing.T) {
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
			name: "test2",
			args: args{
				s: "12:05:45AM",
			},
			want: "19:05:45",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeConversion(tt.args.s); got != tt.want {
				t.Errorf("timeConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	type args struct {
		array []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeConversion1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeConversion(tt.args.s); got != tt.want {
				t.Errorf("timeConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumBribes(t *testing.T) {
	type args struct {
		q []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		// TODO: Add test cases.
		{
			name: "test_2",
			args: args{
				q: []int32{2, 5, 1, 3, 4},
			},
			want: []int32{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumBribes(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumBribes() = %v, want %v", got, tt.want)
			}
		})
	}
}
