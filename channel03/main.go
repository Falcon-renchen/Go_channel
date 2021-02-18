package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int)  {
	for job := range jobs {
		fmt.Printf("worker:%d start job:%d\n",id,job)
		result <- job*2
		time.Sleep(time.Millisecond*500)
		fmt.Printf("worker:%d stop job:%d\n",id,job)
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	//开启三个goroutine
	for i:=0; i<3; i++ {
		go worker(i,jobs,result)
	}
	//发布5个任务
	for i:=0; i<5; i++ {
		jobs <- i
	}
	close(jobs)

	//输出结果
	for i:=0; i<5; i++ {
		res := <- result
		fmt.Println(res)
	}
}
