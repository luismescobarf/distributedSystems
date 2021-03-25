package main

import ("fmt")

func doble(n int) int { return n * 2 }

func main(){
	var ch chan int
	fmt.Printf("%T , %v \n",ch, ch)
}