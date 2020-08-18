package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)
	for i:=0;i<5000;i++{
		go printHelloWorld(i,ch)
	}

	for{
		msg := <- ch
		fmt.Println(msg)
	}

	//time.Sleep(time.Millisecond)
}

func printHelloWorld(i int,ch chan string)  {
	for {
		// 写入 chanl
		ch <- fmt.Sprintf("Hello world from gorouteing %d\n", i)
	}
}