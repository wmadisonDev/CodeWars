package kata

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoOldestAges(t *testing.T) {
	type args struct {
		ages []int
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{
			"1st",
			args{ages: []int{6, 5, 83, 5, 3, 18}},
			[2]int{18, 83},
		},
		{
			"2nd",
			args{ages: []int{1, 5, 87, 45, 8, 8}},
			[2]int{45, 87},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoOldestAges(tt.args.ages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoOldestAges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidParentheses(t *testing.T) {
	type args struct {
		parens string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1st",
			args{parens: "()"},
			true,
		},
		{
			"1st",
			args{parens: ")"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidParentheses(tt.args.parens); got != tt.want {
				t.Errorf("ValidParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveNb(t *testing.T) {
	tests := []struct {
		m    uint64
		want [][2]uint64
	}{
		{
			26,
			[][2]uint64{{15, 21}, {21, 15}},
		},
		{
			100,
			nil,
		},
		{
			101,
			[][2]uint64{{55, 91}, {91, 55}},
		},
		{
			102,
			[][2]uint64{{70, 73}, {73, 70}},
		},
		{
			110,
			[][2]uint64{{70, 85}, {85, 70}},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("RemoveNb(%v)", tt.m), func(t *testing.T) {
			if got := RemoveNb(tt.m); !assert.Equal(t, tt.want, got) {
				t.Errorf("RemoveNb(%v)= %v, want %v", tt.m, got, tt.want)
			}
		})
	}
}