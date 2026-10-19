package main

import (
	"fmt"
	"time"
)

func sayHello1() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello", i)
		time.Sleep(500 * time.Millisecond)
	}
}
func exercise_concurrent1() {
	go sayHello1()
	time.Sleep(800 * time.Millisecond)
	fmt.Println("Exercise9 Finished")
	time.Sleep(2 * time.Second)
}
