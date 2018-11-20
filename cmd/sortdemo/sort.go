package main

import (
	"sort"
	"fmt"
)

/**
*@desc: 实现一个简单排序
*@author:liuhefei
*@date: 
*/

func main(){
	//创建一个数组
	a := []int{8, 5, 0, 1, 3, 7, 4, 2, 9, 6}
	//排序
	sort.Ints(a)

	for i,v := range a  {
		fmt.Println(i, v )
	}

	for _,v := range a  {
		fmt.Println(v)
	}

}