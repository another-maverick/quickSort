package main

import "fmt"


// function that takes an array, low, high and returns the pivot element's sorted position
func Partition(InpArr []int, low, high int) int {

	if high - low < 1 {
		fmt.Println("array length less than 2. Either incorrect input or end of recursion")
	}

	var pivot = InpArr[low]
	var i = low
	var j = high

	for i < j {
		if InpArr[i] <= pivot{
			i++
		}
		if InpArr[j] > pivot{
			j--
		}
		if i < j{
			temp := InpArr[i]
			InpArr[i] = InpArr[j]
			InpArr[j] = temp
		}

	}
	temp := InpArr[low]
	InpArr[low] = InpArr[j]
	InpArr[j] = temp
	return j
}

//function that calls the Partition recursively
func QuickSort(InpArr []int, low, high int) {
	if low < high {
		j := Partition(InpArr, low, high)
		QuickSort(InpArr, low, j)
		QuickSort(InpArr, j+1, high)
	}
}

func main(){
	var InpArr = []int{4,6,3,2,1,7,5,9,8,20,34,66,51}
	fmt.Printf("Before sorting the input array is %v \n", InpArr)
	QuickSort(InpArr, 0,len(InpArr)-1)
	fmt.Printf("After sorting the array is %v \n", InpArr)
}