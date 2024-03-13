package main

import (
	"fmt"

	"github.com/roseduan/minibitcask"
)

var Db *minibitcask.MiniBitcask
var err error

// 定义接受变量
var key, value []byte
var order int

var c = make(chan int)

func main() {
	//将路径修改到本项目文件夹目录
	Db, err = minibitcask.Open("/home/wang2/mini-rosedb-Learn")
	if err != nil {
		panic(err)
	}

	// var (
	// 	key   = []byte("dbname")
	// 	value = []byte("minibitcask")
	// )

	fmt.Println("请输入命令：")
	fmt.Println("1.put")
	fmt.Println("2.get")
	fmt.Println("3.del")
	fmt.Println("4.merge")
	fmt.Println("5.close")

	go func() {
		for {
			fmt.Scanln(&order)
			ReceiveOrder(order)
		}
	}()

	<-c

	fmt.Println("end!")
}

func ReceiveOrder(order int) {
	switch order {
	case 1:
		fmt.Println("请输入key和value，并以 空格 作为分割")
		fmt.Scan(&key, &value)
		err = Db.Put(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Printf("1. put kv successfully, key: %s, value: %s.\n", string(key), string(value))
	case 2:
		fmt.Println("请输入要查询的key")
		fmt.Scan(&key)
		cur, err := Db.Get(key)
		if err != nil {
			panic(err)
		}
		fmt.Printf("get value of key %s, the value of key %s is %s.\n", string(key), string(key), string(cur))

	case 3:
		fmt.Println("请输入要删除的key")
		fmt.Scan(&key)
		err = Db.Del(key)
		if err != nil {
			panic(err)
		}
		fmt.Printf("delete key %s.\n", string(key))

	case 4:
		Db.Merge()
		fmt.Println("compact data to new dbfile.")

	case 5:
		Db.Close()
		c <- 0
	default:
		fmt.Println("指令错误，请输入正确的指令")
		c <- 0
	}
}
