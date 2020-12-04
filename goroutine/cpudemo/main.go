package main

import (
	"fmt"
	"runtime"
)

func main() {
	//获取当前系统cpu的数量
	num := runtime.NumCPU()

	//我们这里设置num的cpu运行go程序
	runtime.GOMAXPROCS(num)
	fmt.Println("num=", num)

}
