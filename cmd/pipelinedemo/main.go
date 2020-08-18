package main

import (
	"bufio"
	"fmt"
	"gointro/pipeline"
	"os"
)

func main() {
	const filename = "small.in"
	const count = 64 // 100 M->000 byte->000
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	p := pipeline.RandomSource(count)
	writer := bufio.NewWriter(f)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	fr, err := os.Open(filename)
	reader := bufio.NewReader(fr)
	p = pipeline.ReadSource(reader, -1)
	i := 0
	for v := range p {
		fmt.Println(v)
		i++
		if i > 10 {
			break
		}
	}

}

func meargeDemo() {
	p := pipeline.Mearge(
		pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(4, 9, 8, 5, 1)))

	for v := range p {
		fmt.Println(v)
	}
}
