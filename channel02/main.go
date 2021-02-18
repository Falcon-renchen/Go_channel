package main

import "fmt"

/**
	协程之间的通讯
	生成0-100个数字存到ch1中，然后从ch1中取出数据平方后传到ch2中
 */
func f1(ch chan<- int)  {
	for i:=0; i<100; i++ {
		ch <- i
	}
	//<-ch 	参数定义ch只能发送数据
	close(ch)
}

func f2(ch1 <-chan int, ch2 chan<- int)  {
	for {
		tmp, ok := <- ch1
		if !ok {
			break
		}
		ch2 <- tmp*tmp
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)

	go f1(ch1)
	go f2(ch1,ch2)

	for ret := range ch2 {
		fmt.Println(ret)
	}
}
