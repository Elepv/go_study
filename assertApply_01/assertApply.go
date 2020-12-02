package main

import "fmt"

//Usb 声明一个接口
type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

//Start 让phone实现Usb的接口方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

func (p Phone) Call() {
	fmt.Println("手机开始打电话...")
}

type Camera struct {
	name string
}

//Start 让Camera实现Usb的接口方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()

	//类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}

	usb.Stop()
}

func main() {
	//定义一个usb接口数组，可以存放Phone和Camera的结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"iPhone"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}

	//遍历usbArr
	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
}
