package main

import "testing"

func Test_titleToNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "FXS",
			},
			want: 4699,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.s); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
