package Decompress_Run_Length_Encoded_List

import (
	"reflect"
	"testing"
)

func Test_decompressRLElist(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				nums: []int{1, 1, 2, 3},
			},
			want: []int{1, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decompressRLElist(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decompressRLElist() = %v, want %v", got, tt.want)
			}
		})
	}
}
