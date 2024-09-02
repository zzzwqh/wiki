package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
)

// 并发写文件 Goroutine

func WriterRole(wg *sync.WaitGroup, writeChan chan interface{}) {
	writeChan <- rand.Intn(100)
	defer wg.Done()
}

func FileReceiver(writeChan chan interface{}, file io.Writer, isDone chan interface{}) {
	for res := range writeChan {
		fmt.Fprintln(file, res)
	}
	isDone <- "Done"
}

func main() {
	var writeChan chan interface{} = make(chan interface{})
	var isDone chan interface{} = make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go WriterRole(&wg, writeChan)
	}
	go func() {
		wg.Wait()
		close(writeChan)
	}()
	file, err := os.OpenFile("goroutine_writeTest.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	go FileReceiver(writeChan, file, isDone)
	<-isDone
}
