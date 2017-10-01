package main

import (
	"fmt"
)
//Mergesort function sort a array pass as a argument
func Mergesort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	mid := len(s) / 2
	l := Mergesort(s[:mid])
	r := Mergesort(s[mid:])
	res := merge(l,r)
	return res 
}

func merge(l,r []int) []int {
	res := make([]int,0,len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(res,r...)
		}
		if len(r) == 0 {
			return append(res,l...)
		}
		if l[0] <= r[0] {
			res = append(res,l[0])
			l = l[1:]
		} else {
			res = append(res,r[0])
			r = r[1:]
		}
	}
	return res
}

func main() {
	s := make([]int,0,2000)
	for i := 2000 ; i > 0 ; i-- {
		s = append(s,i)
	}
	fmt.Printf("sortrd :%v",Mergesort(s))
}