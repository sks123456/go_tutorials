package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main(){
 var intNum int = 32767
 intNum = intNum+1
 fmt.Println(intNum)

 var floatNum float64 = 12345678.9
 fmt.Println(floatNum)

 var floatNum32 float32 = 10.1
 var intNum32 int32 = 2
 var result float32 = floatNum32 + float32(intNum32)
 fmt.Println(result)

 var intNum1 int = 3
 var intNum2 int = 2
 fmt.Println(intNum1 / intNum2)
 fmt.Println(intNum1 % intNum2)

 var myString string = "Hello" + " " + "World"
 fmt.Println(myString)

 fmt.Println(utf8.RuneCountInString("γ"))

 var myRune rune = 'a'
 fmt.Println(myRune) //new data type a type of char

 var myBoolean bool = false
 fmt.Println(myBoolean)

 var intNum3 rune   //go has default value for all integers/uInt/Float, string = '', bool = false
 fmt.Println(intNum3)

 myVar:= "text" // can be used as var myVar = "text"
 fmt.Println(myVar)

 var1, var2 := 1,2 //possble
 fmt.Println(var1,var2)

 const myConst string = "must have value"
 fmt.Println(myConst)

 var printValue string = "Hellow World"
 printMe(printValue)

 var numerator int= 11
var denominator int = 2
var result2, remainder, err = intDivision(numerator, denominator)

// if err!=nil {
// 	fmt.Printf(err.Error())
// }else if remainder == 0 {
// 	fmt.Printf("The result of the integer division is %v ", result2)
// }else{
// 	fmt.Printf("The result of the integer is %v with remainder %v", result2, remainder)
// }


switch{
case err != nil:
	fmt.Printf(err.Error())		//break is automatically applied
case remainder == 0:
	fmt.Printf("The result of the integer division is %v ", result2)
default:
	fmt.Printf("The result of the integer is %v with remainder %v", result2, remainder)
}

switch remainder{
case 0:
	fmt.Printf("The division was exact")
case 1,2:
	fmt.Printf("The division was close")
default:
	fmt.Printf("Not even close")
}
}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int ) (int,int,error) {
	var err error
	if denominator == 0{
		err = errors.New("Cannot divide by zero")
		return 0,0, err
	}

	var result int = numerator/denominator
	var remainder int = numerator%denominator

	return result, remainder, err
}