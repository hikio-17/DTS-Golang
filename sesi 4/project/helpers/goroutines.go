package helpers

import (
	"fmt"
	"runtime"
	"time"
)

func Goroutines() {
	fmt.Println("=============")

	fmt.Println("main execution started")

	go firstProcess(3)

	secondProcess(3)

	fmt.Println("No. of Goroutines :", runtime.NumGoroutine())
	time.Sleep(time.Second * 2)

	fmt.Println("main execution ended")

}

func firstProcess(index int) {
	fmt.Println("First process func started")

	for i := 1; i <= index; i++ {
		fmt.Println("i : ", i)
	}

	fmt.Println("First process func ended")
}

func secondProcess(index int) {
	fmt.Println("Secon process func started")

	for j := 1; j <= index; j++ {
		fmt.Println("j : ", j)
	}

	fmt.Println("Second process func ended")
}
