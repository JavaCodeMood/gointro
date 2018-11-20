package main

import (
	"fmt"
	"time"
)

/**
*@desc: 并发版helloworld
*@author:liuhefei
*@date: 2018/11/11
*/

func main()  {
	//开5个线程
	for i:=0; i< 5;i++{
		//go协程
		go printHelloWorld(i)
	}

	//休眠10毫秒
	time.Sleep(time.Millisecond * 10)
}

func printHelloWorld(i int){
	//死循环，  10毫秒内5个线程抢着输出
	for  {
		fmt.Printf("Hello world! %d！\n", i)
	}

}
