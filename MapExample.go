The Go Playground
 Imports

Hello, playground
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
package main

import (
	"fmt"
)

func main() {
	var numbers []int
	numbers = make([]int, 5, 5)
	fmt.Println("Hello, playground", numbers)
	numbers[0], numbers[1] = 1, 2
	fmt.Println(numbers[0], numbers[1])
	var userMap map[string]string
	userMap = make(map[string]string)
	userMap["France"] = "Paris"
	userMap["Italy"] = "Rome"
	userMap["Japan"] = "Tokyo"
	userMap["India"] = "New Delhi"
	for i := range userMap {
		fmt.Println(i, userMap[i])

	}

}

