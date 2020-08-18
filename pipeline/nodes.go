package pipeline

import (
	"encoding/binary"
	"io"
	"math/rand"
	"sort"
)

// 读取数据到通道
func ArraySource(a ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _,v := range a{
			ch <- v
		}
		close(ch)
	}()
	return ch
}

//排序
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		a := [] int{}
		//read
		for v := range in{
			a = append(a, v)
		}
		//sort
		sort.Ints(a)

		//output
		for _,v := range a{
			out <- v
		}
		close(out)
	}()

	return out
}


// mearge
func Mearge(in1,in2 <-chan int) <-chan int{
	out := make(chan int)
	go func() {
		v1,ok1 := <- in1
		v2,ok2 := <- in2

		for ok1 || ok2 {
			// in2 的内容已经完了
			if !ok2 || (ok1 && v1 <= v2){
				out <- v1
				// 更新v1,ok1的值
				v1,ok1 = <- in1
			}else{
				out <- v2
				v2,ok2 = <- in2
			}
		}
		close(out)
	}()
	return out
}


func ReadSource(reader io.Reader) <- chan int{
	out := make(chan int)
	go func() {
		buffer := make([]byte,8)
		for{
			n,err := reader.Read(buffer)
			if n>0{
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

func WriterSink(writer io.Writer,in <- chan int)  {
	for v := range in{
		buffer := make([]byte,8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

// 生成随机数
func RandomSource(count int) <-chan int{
	out := make(chan int)
	go func() {
		for i:=0;i<count;i++{
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

// 多路两两归并
func MergeN(inputs ... <- chan int,) <- chan int{
	if len(inputs) == 1{
		return inputs[0]
	}
	m := len(inputs) / 2
	return MergeN(
		MergeN(inputs[:m]...),
		MergeN(inputs[m:]...))
}