package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//定义一个全局的pool连接池
var pool *redis.Pool

const ip = "192.168.1.110:6379"

//当启动程序时，就初始化连接池
func init() {

	pool = &redis.Pool{
		MaxIdle:     8,   //表示空闲连接数
		MaxActive:   0,   //表示和数据库的最大连接数，0 表示没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化链接代码，链接哪个ip
			return redis.Dial("tcp", ip)
		},
	}
}

func main() {

	// conn, err := redis.Dial("tcp", ip)
	// if err != nil {
	// 	fmt.Println("connect redis error:", err)
	// 	return
	// }

	//先从pool中取出一个连接
	conn := pool.Get()
	defer conn.Close()
	defer pool.Close()
	fmt.Println("连接redis数据库成功！")

	//通过go向redis写入数据
	_, err := conn.Do("HMSet", "userinf", "uid", "001", "name", "Cris")
	if err != nil {
		fmt.Println("hmset err", err)
		return
	}

	//通过go向redis取出数据。去单个值用redis.String;取多个值为Strings
	res, err := redis.Strings(conn.Do("HMGet", "userinf", "uid", "name"))
	if err != nil {
		fmt.Println("hmget err", err)
	}

	//因为返回的值res 的类型是空接口类型interface，所以使用前必须转成对应类型
	// nameString := res.(string)
	for idx, v := range res {
		fmt.Printf("获得redis数据，res[%d]=%s\n", idx, v)
	}

	fmt.Println("操作数据库OK！")

}
