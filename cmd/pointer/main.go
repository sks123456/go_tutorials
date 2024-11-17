package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v", *p) // pointer value
	fmt.Printf("\nThe value if i is:%v", i)// pointer address
	// *p = 10 //assign value
	p = &i // link the pointer to the i
	*p = 1

	fmt.Printf("\nThe value p points to is :%v", *p)
	fmt.Printf("\nThe value if i is :%v\n", i)	//i change along with p

	var k int32 = 2
	i = k // i just copy the value but not the memory add


	var slice = []int32{1,2,3}		//copying slices is the same as copying each of their respective pointers
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	
}
