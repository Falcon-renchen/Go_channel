package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	for i:=0; i<10; i++ {
		select {
		case x := <- ch:
			fmt.Println(x)
		case ch<-i:
			fmt.Println("存不了数据，",i)
		default:
			fmt.Println("啥也不干")
		}
	}
}
