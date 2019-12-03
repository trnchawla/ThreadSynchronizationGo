package main

import (
	"fmt"
	"time"
)

func printEven(max int, odd chan bool, even chan bool) {
	num := 0
	for num < max {
		fmt.Println(num)
		num += 2
		even <- true
		<-odd
	}
}

func printOdd(max int, even chan bool, odd chan bool) {
	num := 1
	for num < max {
		<-even
		fmt.Println(num)
		num += 2
		odd <- true
	}
}
func main() {
	var evenChan chan bool = make(chan bool)
	var oddChan chan bool = make(chan bool)
	go printEven(100, oddChan, evenChan)
	go printOdd(100, evenChan, oddChan)
	time.Sleep(3 * time.Second)
}
