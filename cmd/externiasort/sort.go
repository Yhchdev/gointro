package main

import (
	"bufio"
	"fmt"
	"gointro/pipeline"
	"os"
)

func main() {
	p := createPipeline(
		"small.in", 512, 4)
	writeToFile(p, "")
	printFile("small.out")

}

func printFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	source := pipeline.ReadSource(reader, -1)

	for v := range source {
		fmt.Println(v)
	}

}

func writeToFile(p <-chan int, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)

}

func createPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount

	sortResults := []<-chan int{}

	for i := 0; i < chunkCount; i++ {
		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		f.Seek(int64(chunkSize*i), 0)
		source := pipeline.ReadSource(bufio.NewReader(f), chunkSize)
		sortResults = append(sortResults, pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResults...)
}
