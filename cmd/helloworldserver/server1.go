package main

import (
	"net/http"
	"fmt"
)

/**
*@desc: http://localhost:8081/?name=liuhefei
*@author:liuhefei
*@date: 2018/11/11
*/
func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//传递一个参数
		fmt.Fprintf(writer,"<h1 style='color:blue'>Hello, %s !</h1>", request.FormValue("name"))
	})

	http.ListenAndServe(":8081", nil)
}
