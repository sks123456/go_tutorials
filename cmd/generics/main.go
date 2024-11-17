package main

import "fmt"

func main(){
	var intSLice = []int{}
	fmt.Println(sumSlice[int](intSLice))
	fmt.Println(isEmpty((intSLice)))

	var float32Slice = []float32{1,2,3}
	fmt.Println(sumSlice[float32](float32Slice))
	fmt.Println(isEmpty((float32Slice)))

}

func isEmpty [T any](slice []T) bool{
	return len(slice)==0
}

func sumSlice[T int | float32| float64](slice []T) T{
	var sum T
	for _,v := range slice{
		sum += v
	}
	return sum
}