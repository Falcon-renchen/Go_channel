package main

import "fmt"

func main() {
/*	var ch1 chan int
	ch1 = make(chan int, 1)*/

	//ch1 := make(chan int) 	//无缓冲区的通道,同步通道
	ch1 := make(chan int, 1)	//channel 是一个引用值		带缓冲通道，异步通道
	ch1 <- 10
	x := <- ch1
	//len(ch1)
	//cap(ch1) //通道的容量
	fmt.Println(x)
	close(ch1)
}
