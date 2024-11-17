package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main(){
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"wallmart.com","costco.com","wholefood.com"}
	for i:= range websites{
		go checkChickenPrices(websites[i],chickenChannel)
		go checkTofuPrices(websites[i],tofuChannel)
	}
	sendMessage(chickenChannel,tofuChannel)
}

func checkTofuPrices(website string, c chan string){
	for {
		time.Sleep(time.Second*1)
		var tofu_price = rand.Float32()*20
		if tofu_price<=MAX_CHICKEN_PRICE{
			c <- website
			break
		}
	}
}

func checkChickenPrices(website string, chickenChannel chan string){
	for {
		time.Sleep(time.Second*1)
		var chickenPrice = rand.Float32()*20
		if chickenPrice<=MAX_CHICKEN_PRICE{
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string){
	select{
	case website:= <-chickenChannel:
		fmt.Printf("\nText sent: Found a deal on chicken at %v",website)
	case website := <-tofuChannel:
		fmt.Printf("\nText sent: Found a deal on tofu at %v",website)
	}
}