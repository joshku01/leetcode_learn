package Two_Out_of_Three

import (
	"reflect"
	"testing"
)

func Test_twoOutOfThree(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
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
				nums1: []int{1, 1, 3, 2},
				nums2: []int{2, 3},
				nums3: []int{3},
			},
			want: []int{3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoOutOfThree(tt.args.nums1, tt.args.nums2, tt.args.nums3); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoOutOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
