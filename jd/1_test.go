package main

import (
	"fmt"
	"testing"
)

func Test_isSurround(t *testing.T) {
	type args struct {
		nums [][]rune
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				nums: [][]rune{
					{'o', 'o', 'o'},
					{'.', '.', '.'},
					{'.', '.', '.'},
				},
				x: 1, y: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSurround(tt.args.nums, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isSurround() = %v, want %v", got, tt.want)
			}
			fmt.Println(true)
		})
	}
}
