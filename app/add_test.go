package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_number_add(t *testing.T) {
	type fields struct {
		n int
	}
	type args struct {
		otherNumber number
	}
	tests := []struct {
		name   string
		n int
		otherNumber number
		want   number
	}{
		{
			name: "(+) Should return the addition of two number",
			n : 4,
			otherNumber: number{
				n: 6,
			},
			want: number{n: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := number{
				n: tt.n,
			}
			assert.Equal(t, tt.want, num.add(tt.otherNumber))
		})
	}
}
