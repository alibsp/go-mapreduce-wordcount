package main

import (
	"fmt"
	"sync"
	"time"
)

/*
من 3 کار دارم.

        ↓

کار 1 ─────┐
کار 2 ─────┼──→ منتظر بمان
کار 3 ─────┘

        ↓

هر سه تمام شدند

        ↓

ادامه بده
*/

func sayHello2(id int, wg *sync.WaitGroup) {
	defer wg.Done() // اعلام پایان کار یکی از روتین ها- به عبارتی یکی از شمارنده کم می شود
	//defer: یعنی این دستور را هنگام خروج از تابع اجرا کن
	fmt.Println("Hello from", id)
	if id == 1 {
		time.Sleep(2 * time.Second)
		fmt.Println("finish sleep at", id)
	}
}
func exercise_concurrent2() {
	var wg sync.WaitGroup
	wg.Add(3) // اعلام تعداد کارها
	go sayHello2(1, &wg)
	go sayHello2(2, &wg)
	go sayHello2(3, &wg)
	wg.Wait() //انتظار برای پایان همه روتین ها
	fmt.Println("All goroutines finished")
}
