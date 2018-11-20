package main

import (
	"os"
	"gointro/pipeline"
	"bufio"
	"fmt"
)

/**
*@desc: 高并发并行计算外部排序
*@author:liuhefei
*@date: 2018/11/11
*/
func main() {
	//1.生成pipeline
	//p := createPipeline("small.in", 512, 4)  //小数据, 4:表示4个节点
	//2.写入文件
	//writeToFile(p, "small.out")
	//3.打印文件
	//printFile("small.out")

	p := createPipeline("large.in", 800000000, 4)   //大数据
	//2.写入文件
	writeToFile(p, "large.out")
	//3.打印文件
	printFile("large.out")
}

//单机版
func createPipeline(filename string, filesize, chunkCount int) <- chan int{
	//获取分块的大小
	chunkSize := filesize / chunkCount
	pipeline.Init()
	//收集结果
	sortResults := [] <- chan int{}
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
		sortResults = append(sortResults, pipeline.InMemSort(source))
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