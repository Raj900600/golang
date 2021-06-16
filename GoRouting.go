package main

import (
    "fmt"
    "time"
)

func main(){
	fmt.Println("Go start")
msg:="hello"
go func (m string)  {
	fmt.Println(m)
}(msg);
msg = "HelloWorld"
	fmt.Println("Go End")
	time.Sleep(8 * time.Second)
}
func hello(grn int)  {
	for i:=0;i<5;i++{
		fmt.Println(grn,":Hello")
	}
	
}