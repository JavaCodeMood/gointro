package main

import (
	"net/http"
	"fmt"
)

/**
http://localhost:8080/
go语言中：参数名在类型的前面
go语言的函数可以作为参数传递
*@desc: http://localhost:8080/
*@author:liuhefei
*@date: 2018/11/11
*/
func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "<h1 style='color:red'>Hello world!</h1>")
	})

	http.ListenAndServe(":8080",nil)
}
