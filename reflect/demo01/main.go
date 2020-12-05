package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {

	//通过反射获取传入的变量的type， kind， 值
	//先获取reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)
	fmt.Printf("rVal=%v; r/val=%T\n", rVal, rVal)
}

func main() {

	//基本数据类型的reflect机制
	var num int = 100
	reflectTest(num)

}
