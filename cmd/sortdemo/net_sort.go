package main

import (
	"gointro/pipeline"
	"os"
	"bufio"
	"fmt"
	"strconv"
)

/**
*@desc: 网络版高并发排序
*@author:liuhefei
*@date: 2018/11/12 9:52
*/
func main() {
	//1.生成pipeline
	p := createNetworkPipeline("small.in", 512, 4)  //小数据, 4:表示4个节点
	//2.写入文件
	writeToFile(p, "small.out")
	//3.打印文件
	printFile("small.out")

	//p := createNetworkPipeline("large.in", 800000000, 4)   //大数据
	//time.Sleep(time.Hour)
	//2.写入文件
	//writeToFile(p, "large.out")
	//3.打印文件
	//printFile("large.out")
}

//网络版
func createNetworkPipeline(filename string, filesize, chunkCount int) <- chan int{
	//获取分块的大小
	chunkSize := filesize / chunkCount
	pipeline.Init()
	//收集结果
	sortAddr := []string{}
	for i := 0;i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil{
			panic(err)
		}

		//如果filesize / chunkCount不整除
		//i:表示第几块，从第几块开始； 0：表示从头开始
		file.Seek(int64(i * chunkSize), 0)

		//读取数据
		source := pipeline.ReaderSourceOne(bufio.NewReader(file), chunkSize)
		//收集结果
		addr := ":" + strconv.Itoa(7001 + i)
		pipeline.NetWorkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}

	sortResults := []<-chan int{}
	for _,addr := range sortAddr {
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sortResults...)
}

//将数据写入文件
func writeToFile(p <- chan int, filename string){
	//创建文件
	file,err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//写入文件
	writer := bufio.NewWriter(file)
	//刷新
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

//读取文件的内容
func printFile(filename string){
	//打开文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//读取数据
	p := pipeline.ReaderSourceOne(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count ++
		if count >= 100 {   //做限制
			break
		}
	}

}


