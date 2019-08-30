package main

import "fmt"

func merge(l, r []int) []int {
	finalArray := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		fmt.Println("Array is", l, r)
		if len(l) == 0 {
			fmt.Println("len is ", len(l), "append", r)
			return append(finalArray, r...)
		}
		if len(r) == 0 {
			fmt.Println("len is ", len(r), "append", l)
			return append(finalArray, l...)
		}
		if l[0] <= r[0] {
			finalArray = append(finalArray, l[0])
			fmt.Println(r[0], "is bigger than", l[0], "append", l[0])
			l = l[1:]
		} else {
			finalArray = append(finalArray, r[0])
			fmt.Println(l[0], "is bigger than", r[0], "append", r[0])
			r = r[1:]
		}
	}
	return finalArray
}

func mergeSort(randomArray []int) []int { // initialise empty array
	if len(randomArray) <= 1 {
		return randomArray
	}
	dividedArray := len(randomArray) / 2               // divide array size
	leftArray := mergeSort(randomArray[:dividedArray]) //recursive dividing of the arrays
	rightArray := mergeSort(randomArray[dividedArray:])
	return merge(leftArray, rightArray)
}

func main() {
	randomArray := []int{5, 4, 1, 8, 7, 2, 6, 3}
	fmt.Printf("%v\n%v\n", randomArray, mergeSort(randomArray))
}
