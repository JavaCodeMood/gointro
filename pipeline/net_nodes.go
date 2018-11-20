package pipeline

import (
	"net"
	"bufio"
)

/**
*@desc: 网络版高并发并行计算排序工具
*排序原理：1.分块读取数据，每一块内部做排序
*2.对每个内部排序的节点都做了一个server，用于监听客户的连接，客户一旦连接成功，就把排好序的数据传送给它
*3.对每一块排好序的块做两两归并，在对两块归并好的归并块做归并，最后得出排序结果
*@author:liuhefei
*@date: 2018/11/12
*/

//往网络写
func NetWorkSink(addr string, in <- chan int){
	//监听端口
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	go func() {
		defer listener.Close()

		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		//写数据
		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		WriterSink(writer, in)
	}()

}

func NetworkSource(addr string) <- chan int {
	out := make(chan int)
	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil{
			panic(err)
		}
		r := ReaderSourceOne(bufio.NewReader(conn), -1)
		for v := range r {
			out <- v
		}
		close(out)
	}()
	return out
}
