package mylib

import (
	"fmt"
	"time"
)

var c chan int

func Ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready！")
	c <- 1
}

func RunReady() {
	var i int
	c = make(chan int)
	go Ready("Tea", 2)
	go Ready("Coffee", 1)
	fmt.Println("I'am waiting,but not to long")
	//time.Sleep(2 * time.Second)
	//<-c
	//<-c
L:
	for {
		select {
		case <-c:
			i++
			if i > 1 {
				break L
			}
		}
	}

}
