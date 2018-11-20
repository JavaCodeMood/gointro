package main

import (
	"gointro/pipeline"
	"fmt"
	"os"
	"bufio"
)

/**
*实现大数据归并排序
*
*@desc: 归并排序工具
*@author:liuhefei
*@date: 2018/11/11
*/
func main() {
	//小文件
	//const filename  = "small.in"
	//const n  = 50

	//大数据
	const filename  = "large.in"
	const n  = 100000000   //将会生成 100000000kb * 8 = 800MB 大小的数据

	//创建文件
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
    defer file.Close()

	//生成50个随机数
	p := pipeline.RandomSource(n)

	writer := bufio.NewWriter(file)
	//将生成的随机数写入文件
	pipeline.WriterSink(writer, p)
	writer.Flush()

	//读取文件中的数据
	file,err = os.Open(filename)  //打开文件
    if err != nil {
    	panic(err)
	}
	defer file.Close()

	//p = pipeline.ReaderSource(bufio.NewReader(file))  //读取数据
	p = pipeline.ReaderSourceOne(bufio.NewReader(file), -1)  //分块读取数据
	count := 0
	for v := range p {
		fmt.Println(v)
		count ++
		if(count >= 100){   //只打印前100个数据
			break
		}
	}

}

func mergeDemo(){
	p := pipeline.Merge(pipeline.InMemSort(pipeline.ArraySourceOne(7, 0, 1, 4, 6, 9, 2)),
		pipeline.InMemSort(pipeline.ArraySourceOne(10,3,5,8,7,11,24,13)))
	for  {
		if num, ok := <- p; ok{
			fmt.Println(num)
		}else{
			break
		}
	}
	fmt.Println("----------------------")
	for v := range p{
		fmt.Println(v)
	}

}