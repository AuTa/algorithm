package leetcode0334

import "testing"

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "12345",
			args: args{
				[]int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "54321",
			args: args{
				[]int{5, 4, 3, 2, 1},
			},
			want: false,
		},
		{
			name: "25345",
			args: args{
				[]int{2, 5, 3, 4, 5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
