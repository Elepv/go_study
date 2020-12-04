package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {

	for i := 1; i <= 500000; i++ {
		intChan <- i
	}

	//关闭intChan
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	//使用for循环
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true //假设是素数
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			//将这个数放入到primeChan
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	//协程取不到数据，则向exitChan写入一个true
	exitChan <- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	//标识退出的管道
	exitChan := make(chan bool, 4)

	fmt.Println("程序开始统计素数")
	start := time.Now().Unix()

	//开启一个协程，想intChan放入1-8000个数
	go putNum(intChan)

	//开启4个协程，从intChan取出数据，并判断是否为素数，如果是，就放到管道primeChan中去
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		//统计使用协程所需要的时间
		end := time.Now().Unix()
		fmt.Println("使用协程耗时：", end-start)

		//当我们从exitChan取出了4个结果，就看可以关闭primeChan
		close(primeChan)
	}()

	//里边我们的primeNum，把结果取出来
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		// fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main线程退出")
}
