package main

import (
	"testing"
	)
func TestHandler(t *testing.T) {
	cases := []struct {
		in, out []int
	}{
		{
			[]int{8,7,-2},[]int{-2,7,8},
		},
		{
			[]int{5,0,9},[]int{0,5,9},
		},
		{
			[]int{9,0,-100,-4},[]int{-100,-4,0,9},
		},
	}

	for _,c := range cases {
		var equal = true
		sorted := Mergesort(c.in)
		for i := range c.in {
			if sorted[i] != c.out[i] {
				equal = false
			}
		}
		if !equal {
			t.Errorf("Expected result is %v ,instead get %v",c.out,sorted)
		}
	}
}

func BenchmarkHandler(b *testing.B) {
	var equal = true
	s := []int{2,9,3,7,-8,-64,-415}
	for i := 0 ; i < b.N ; i++ {
		sorted := Mergesort(s)
		expected := []int{-415,-64,-8,2,3,7,9}
		for i := range s {
			if sorted[i] != expected[i] {
				equal = false
			}
		}
		if !equal {
			b.Errorf("Expected result is %v ,instead get %v",expected,sorted)
		}
	}
}