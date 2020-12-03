package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体
type Monster struct {
	Name     string `json:"name"`
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func TestStruct() {
	//演示
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.00,
		Skill:    "牛魔拳",
	}

	//将monster 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("monster 序列化后=%+v\n", string(data))
}

//将map进行序列化
func TestMap() {
	//定义一个map
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "孙悟空"
	m["age"] = 500
	m["address"] = "花果山"

	//将m这个map进行序列化
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("m map 序列化后=%v\n", string(data))
}

func main() {
	//将结构体，map，切片进行序列化
	TestStruct()
	TestMap()
}
