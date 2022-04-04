package myalgo

import "testing"

func Test_IsPrime(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   bool
	}{
		{
			name:   "素数",
			number: 53,
			want:   true,
		},
		{
			name:   "合成数",
			number: 77,
			want:   false,
		},
		{
			name:   "0は素数ではない",
			number: 0,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.number); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
