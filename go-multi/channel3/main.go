package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	if <-exitChan {
		return
	}
	fmt.Println()
}

func writeData(intChan chan int) {
	rand.Seed(time.Now().UnixNano())

	var tempInt int
	for i:=1; i<10; i++ {
		tempInt = rand.Intn(4) + 10
		fmt.Printf("写入到%v,第%v次写\n", tempInt, i)
		intChan<- tempInt
		time.Sleep(time.Second * 1)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	var count int
	for {
		val, ok := <-intChan
		count++
		if !ok {
			break
		}
		fmt.Printf("读取到%v,第%v次读取\n", val, count)
	}
	exitChan <- true
	close(exitChan)
}