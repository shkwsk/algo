package myalgo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestIsPrime(t *testing.T) {
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

func TestDivisorEnumeration(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0は約数なし",
			args: args{
				n: 0,
			},
			want: []int{},
		},
		{
			name: "正常",
			args: args{
				n: 100,
			},
			want: []int{1, 2, 4, 5, 10, 20, 25, 50, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := []cmp.Option{
				cmpopts.SortSlices(func(i, j int) bool {
					return i < j
				}),
			}
			got := DivisorEnumeration(tt.args.n)
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("DivisorEnumeration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrimeFactrization(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "正常",
			args: args{
				n: 100,
			},
			want: []int{2, 2, 5, 5},
		},
		{
			name: "正常",
			args: args{
				n: 286,
			},
			want: []int{2, 11, 13},
		},
		{
			name: "正常",
			args: args{
				n: 20211225,
			},
			want: []int{3, 5, 5, 31, 8693},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := []cmp.Option{
				cmpopts.SortSlices(func(i, j int) bool {
					return i < j
				}),
			}
			got := PrimeFactrization(tt.args.n)
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("PrimeFactrization() = %v, want %v", got, tt.want)
			}
		})
	}
}
