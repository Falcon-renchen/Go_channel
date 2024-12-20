package main

import (
	"fmt"
	"sync"
)

type Cmd func(list []int) chan int
type PipeCmd func(in chan int) chan int

func Multiply(list []int) (ret []int) {
	ret = make([]int, 0)
	for _, num := range list {
		ret = append(ret, num*10)
	}
	return
}

func Pipe(args []int, c1 Cmd, cs ...PipeCmd) chan int {
	ret := c1(args)
	retlist := make([]chan int, 0)

	for index, c := range cs {
		if index == 0 {
			retlist = append(retlist, c(ret))
		} else {
			getChan := retlist[len(retlist)-1]
			retlist = append(retlist, c(getChan))
		}
	}
	return retlist[len(retlist)-1]
}

//多路复用,两个两个输出
func Pipe2(args []int, c1 Cmd, cs ...PipeCmd) chan int {
	ret := c1(args) //找偶数
	out := make(chan int)
	wg := sync.WaitGroup{}

	for _, c := range cs {
		getChan := c(ret)
		wg.Add(1)
		go func(input chan int) {
			defer wg.Done()
			for v := range input {
				out <- v
			}
		}(getChan)
	}
	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}

func Evens(list []int) chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for _, num := range list {
			if num%2 == 0 {
				c <- num
			}
		}
	}()
	return c
}

func M2(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num * 2
		}
	}()

	return out
}

func M10(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num * 10
		}
	}()

	return out
}

func main() {
	nums := []int{2, 3, 5, 12, 22, 16, 4, 9, 23, 64, 62}

	ret := Pipe2(nums, Evens, M10, M2, M10, M2)
	for r := range ret {
		fmt.Printf("%d ", r)
	}
	//fmt.Println(Multiply(Evens(nums)))

	//fmt.Println(p(nums, Evens, Multiply))

}
