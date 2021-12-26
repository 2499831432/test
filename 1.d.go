package main

func main() {
	sieve()
}

//生成一个存有一到一百万的管道，作为原始数据
func generate(ch chan<- int) {
	for i := 1; i <= 1000000; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

//接受数据和发送数据的管道以及此次作为筛选基准的数字
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // 便利初始数据
		if i%prime != 0 {
			dst <- i // 把不能被基准数字整除的数字发送到dst管道中
		}
	}
}

func sieve() {
	ch := make(chan int) // 创建一个管道，后续用于发送数据
	go generate(ch)      // 给管道输入数据
	for i := 2; i <= 1000; i++ {
		ch1 := make(chan int)
		go filter(ch, ch1, i)
		ch = ch1//把筛选过的数据管道赋值给初始管道，最终筛选出的素数也都保存在ch管道中
	}
}
