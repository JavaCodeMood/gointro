package main

import (
	"fmt"
)

/**
*@desc: 并发版helloworld, 通过通道传递
*@author:liuhefei
*@date: 2018/11/11
*/

func main(){
	ch := make(chan string)  //创建一个通道
	for i:=0;i<5;i++{
		//go协程
		go printHelloWorld(i, ch)
	}

	for  {
		msg := <- ch  //通道中的数据传递给msg
		fmt.Println(msg)
	}
}

//使用通道传递参数
func printHelloWorld(i int, ch chan string){
	for  {
		ch <- fmt.Sprintf("Hello world - %d!\n", i)
	}
}
