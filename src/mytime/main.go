//Time package
package main

import (
	
	"fmt"
	"time"
)

func main() {
	
	var t time.Time
	
	//t = time.Now()
	t = time.Date(2018,time.March,11,12,12,12,0,time.Local)
	
	fmt.Println(t.Year())
	fmt.Println(t.YearDay())
	fmt.Println(t.Month())
	fmt.Println(t.Weekday())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Nanosecond())
	fmt.Println(t.Zone())
	
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)
	
	d, _ := time.ParseDuration("2h30m")
	
	fmt.Println(d)
	
	t = time.Now()
	
	t = t.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t)
	
	t = time.Date(2020,time.July,24,0,0,0,0,time.Local)
	
	fmt.Println(t.Sub(time.Now()))
	
	var ch <-chan time.Time
//	ch = time.Tick(3 * time.Second)
//	
//	for {
//		
//		fmt.Println(<-ch)
//	}
	
	//ここで５秒間ゴルーチンは停止する。
	ch = time.After(5 * time.Second)
	fmt.Println(<-ch)
	
	
}

