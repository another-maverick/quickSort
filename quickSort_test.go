package main

import (
	"sort"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {
	for i:=0; i<b.N; i++{
		var NewInp = []int{4,6,3,2,1,7,5,9,8,20,34,66,51, 123, 124,130, 200, 190, 185, 180, 165, 160, 155, 157, 159, 421, 232, 343, 567,32}
		QuickSort(NewInp, 0, len(NewInp)-1)
	}
}

func BenchmarkDefaultSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var NewInp = []int{4,6,3,2,1,7,5,9,8,20,34,66,51, 123, 124,130, 200, 190, 185, 180, 165, 160, 155, 157, 159, 421, 232, 343, 567,32}
		sort.Ints(NewInp)
	}
}