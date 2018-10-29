package main

import (
	"sync"
	"fmt"
	"sync/atomic"
	"time"
)


func addNum(numP *int32, id int, deferFunc func()) {
	defer func(){
		deferFunc()
	}()

	for i:=0;;i++ {
		currNum := atomic.LoadInt32(numP)
		newNum := currNum + 1
		time.Sleep(time.Microsecond * 200)
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			fmt.Printf("The number:%d [%d-%d]\n", newNum, id, i)
			break
		} else {
			//
		}
	}
}

func corrdinateWithWaitGroup() {
	total := 12
	stride := 3
	var num int32
	fmt.Printf("The number:%d [with sync.WaitGroup]\n", num)
	var wg sync.WaitGroup
	for i:= 1;i<=total;i = i + stride {
		wg.Add(stride)
		for j:=0;j<stride;j++ {
			go addNum(&num, i+j, wg.Done)
		}
		wg.Wait()
	}
	fmt.Println("End.")

}


func main() {
	corrdinateWithWaitGroup()
}