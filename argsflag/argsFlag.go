package main

import (
	"flag"
	"fmt"
)

func main() {
	//定义几个变量，用来接收命令行参数
	var user string
	var pwd string
	var host string
	var port int

	//接收命令行参数，&user接收命令行中带-u的参数
	//"u",就是 -u 指定的参数
	//""，为默认值
	//"用户名，默认值为空",说明语句
	flag.StringVar(&user, "u", "", "用户名，默认值为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认值为空")
	flag.StringVar(&host, "h", "localhost", "主机ip，默认值为localhost")
	flag.IntVar(&port, "p", 3306, "端口号，默认值为3306")

	flag.Parse()

	fmt.Printf("用户名=%v, 密码=%v， 主机地址=%v， 端口号=%v", user, pwd, host, port)
}
