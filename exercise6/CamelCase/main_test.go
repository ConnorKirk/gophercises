package main

import "testing"

func Test_countCamelCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{"alphaBeta"},
			want: 2,
		},
		{
			name: "t2",
			args: args{"alpha"},
			want: 1,
		},
		{
			name: "t3",
			args: args{"alphaBetaGammaGGamma"},
			want: 5,
		},

		{
			name: "t4",
			args: args{""},
			want: 0,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCamelCase(tt.args.input); got != tt.want {
				t.Errorf("countCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}

}
