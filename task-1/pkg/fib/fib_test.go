package fib

import (
	"testing"
)

func TestNum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				n: 1,
			},
			want: 0,
		},
		{
			name: "Case 2",
			args: args{
				n: 2,
			},
			want: 1,
		}, {
			name: "Case 10",
			args: args{
				n: 10,
			},
			want: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Num(tt.args.n); got != tt.want {
				t.Errorf("Num() = %v, want %v", got, tt.want)
			}
		})
	}
}
