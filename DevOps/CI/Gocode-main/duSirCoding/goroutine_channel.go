package main

import "fmt"

func writeData(ch chan<- int) {
	for i := 0; i < 50; i++ {
		ch <- i
	}
	close(ch)
}

func readData(ch <-chan int, exitChan chan<- bool) {
	//for res := range ch {
	//	fmt.Println(res)
	//}
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(val)
	}
	exitChan <- true
	close(exitChan) // 因为 close 了这个 channel ，所以 main 函数中无论读多少次都不会堵塞
}

func main() {
	var chIns1 = make(chan int)
	var exitChan = make(chan bool)
	go writeData(chIns1)
	go readData(chIns1, exitChan)

	for {
		exitSignal, ok := <-exitChan
		fmt.Println(exitSignal, ok)

		if ok {
			break
		}
	}
}
