package main

import "fmt"

func main() {
	ch := make(chan int,10)
	for i := 0; i < 10; i++ {
	go func() {
		fmt.Println("在main函数中开启10个协程输出一段话，要求10行话全部输出完毕后再结束main函数。")
		ch <- 2
	}()
	}
	for i:=0;i<10;i++ {
	<-ch
	}
}

//func test() {
//	fmt.Println("在main函数中开启10个协程输出一段话，要求10行话全部输出完毕后再结束main函数。")
//	<-ch
//}
