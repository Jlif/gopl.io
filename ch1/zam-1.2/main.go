package main

import (
	"os"
	"strconv"
	"fmt"
)

//修改echo程序，使其打印value和index，每个value和index显示一行
func main() {
	s, tail := "", ""
	for index, value := range os.Environ()[1:] {
		tail = "\n"
		s += "[" + strconv.Itoa(index) + "]" + value + tail
	}
	fmt.Println(s)
}
