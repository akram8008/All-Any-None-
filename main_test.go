package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	tests := []struct {
		name string
		ls []int
		predicate func(int)bool
		want bool
	}{
		{
			"The value {7} is equal to all items",
			[]int{7,7,7,7},
			func (i int)bool {
				if i==7{
					return true
				}
				return false
			},true},
		{
			"The value {7} is not equal to all items ",
			[]int{1,-5,6,7},
			func (i int)bool {
				if i==7{
					return true
				}
				return false
			},false},

	}

	for _, tt := range tests {
		if got := All(tt.ls, tt.predicate); got != tt.want {
			t.Errorf("Index() = %v, want %v", got, tt.want)
		}
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		name string
		ls []int
		predicate func(int)bool
		want bool
	}{
		{
			"The value {7} exists in the position 3 (order is starting from 0)",
			[]int{1,-5,6,7},
			func (i int)bool {
				if i==7{
					return true
				}
				return false
			},true},
		{
			"The value {7} does not exist in the slice ",
			[]int{1,-5,6},
			func (i int)bool {
				if i==7{
					return true
				}
				return false
			},false},

	}

	for _, tt := range tests {
		if got := Any(tt.ls, tt.predicate); got != tt.want {
			t.Errorf("Index() = %v, want %v", got, tt.want)
		}
	}
}

func TestFind(t *testing.T) {
	var tests = []struct {
		name      string
		ls        []int
		predicate func(int) bool
		want      int
	}{
		{
			"The value {7} exists in the position 3 (order is starting from 0)",
			[]int{1, -5, 6, 7},
			func(i int) bool {
				if i == 7 {
					return true
				}
				return false
			}, 7},
		{
			"The value {7} does not exist in the slice ",
			[]int{1, -5, 6},
			func(i int) bool {
				if i == 7 {
					return true
				}
				return false
			}, -1,
	}}
	for _, tt := range tests {

		if got,err := Find(tt.ls, tt.predicate); got!=tt.want   {
			t.Errorf("Index() = %v,Error=%v, want %v", got,err,tt.want)
		}
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		name string
		ls []int
		predicate func(int)bool
		want int
	}{
	  {
	  "The value {7} exists in the position 3 (order is starting from 0)",
	  []int{1,-5,6,7},
	  func (i int)bool {
	  	if i==7{
	  		return true
		}
		return false
	  },3},
		{
			"The value {7} does not exist in the slice ",
			[]int{1,-5,6},
			func (i int)bool {
				if i==7{
					return true
				}
				return false
			},-1},

	}

	for _, tt := range tests {
		if got := Index(tt.ls, tt.predicate); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
	}
}

func TestNone(t *testing.T) {
	tests := []struct {
		name string
		ls []int
		predicate func(int)bool
		want bool
	}{
		{
			"The value {7} exists in the position 3 (order is starting from 0)",
			[]int{1,-5,6,7},
			func (i int)bool {
				if i==7{
					return true
				}
				return false
			},false},
		{
			"The value {7} does not exist in the slice ",
			[]int{1,-5,6},
			func (i int)bool {
				if i==7{
					return true
				}
				return false
			},true},

	}

	for _, tt := range tests {
		if got := None(tt.ls, tt.predicate); got != tt.want {
			t.Errorf("Index() = %v, want %v", got, tt.want)
		}
	}
}