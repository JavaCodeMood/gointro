package pipeline

import (
	"sort"
	"io"
	"encoding/binary"
	"math/rand"
	"time"
	"fmt"
)

/**
*@desc: 单机版  归并排序工具
*@author:liuhefei
*@date: 2018/11/11
*/

var startTime time.Time

func Init(){
	startTime = time.Now()
}
func ArraySource(a ...int) chan int{
	out := make(chan int)  //创建一个通道
	go func(){
		for _,v := range a{
			out <- v
		}
		close(out)  //关闭
	}()
	return out
}

func ArraySourceOne(a ...int) <- chan int{
	out := make(chan int)  //创建一个通道
	go func(){
		for _,v := range a{
			out <- v
		}
		close(out)  //关闭
	}()
	return out
}

//实现内部排序功能
func InMemSort(in <-chan int) <- chan int{
	out := make(chan int, 1024)
	go func(){
		//读取
		a := []int{}
		for v := range in {
			a = append(a,v)
		}
		fmt.Println("Read done:", time.Now().Sub(startTime))  //当前时间-开始时间
		//排序
		sort.Ints(a)
		fmt.Println("InMemSort done:", time.Now().Sub(startTime))  //当前时间-开始时间

		//输出
		for _, v := range a{
			out <- v
		}
		close(out)
	}()
	return out
}

//两两合并
func Merge(in1, in2 <- chan int) <- chan int{
	out := make(chan int, 1024)   //提高效率
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2){
				out <- v1
				v1, ok1 = <- in1
			}else{
				out <- v2
				v2, ok2 = <- in2
			}
		}
		close(out)
		fmt.Println("Merge done:", time.Now().Sub(startTime))  //当前时间-开始时间
	}()
	return out
}

//从文件中读取数据
func ReaderSource(reader io.Reader) <- chan int{
	out := make(chan int)
	go func() {
		buffer := make([]byte, 8)  //8字节
		for  {
			//n:表示读了多少字节， err:表示是否有错误
			n, err :=reader.Read(buffer)
			if n>0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil{
				break
			}
		}
        close(out)
	}()
	return out
}

//从文件中读取数据,分块读取
func ReaderSourceOne(reader io.Reader, chunkSize int) <- chan int{
	out := make(chan int, 1024)   //发1024个在收
	go func() {
		buffer := make([]byte, 8)  //8字节
		bytesRead := 0
		for  {
			//n:表示读了多少字节， err:表示是否有错误
			n, err :=reader.Read(buffer)
			bytesRead += n
			if n>0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil || (chunkSize != -1 && bytesRead >= chunkSize){
				break
			}
		}
		close(out)
	}()
	return out
}

//写数据到文件
func WriterSink(writer io.Writer, in <- chan int){
	for v := range in{
		buffer := make([]byte, 8)  //创建一个8字节的数组
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

//生成随机数
func RandomSource(count int) <-chan int{
	out := make(chan int)
	go func(){
		for i:=0; i<count; i++{
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

//n个数据两两归并,递归调用
func MergeN(inputs ... <- chan int) <- chan int{
	if len(inputs) == 1{
		return inputs[0]
	}

	m := len(inputs) / 2
	//合并inputs[0..m) and inputs[m..end)
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))

}