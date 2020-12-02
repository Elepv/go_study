package main

import "fmt"

//编写一个函数，可以判断输入的类型

func TypeJudge(items ...interface{}) {

	for idx, x := range items {
		index := idx + 1
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是 float32 类型，值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数是 float64 类型，值是%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是 整数 类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是 string 类型，值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数是 不确定 类型，值是%v\n", index, x)
		}
	}
}

type Student struct{}

func main() {

	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "Tom"
	address := "北京"
	n4 := 300

	stu1 := Student{}
	stu2 := &Student{}

	fmt.Printf("%v\n", stu1)
	fmt.Printf("%v\n", stu2)
	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2)
}
