package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [] int{3,9,1,8,2,7,4,6}
	sort.Ints(a)

	for _,v := range a{
		fmt.Println(v)
	}
}
