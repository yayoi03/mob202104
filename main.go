package main

import "fmt"

type Reservation struct{
	hour,min,minutes int
}

func main() {

	t := Reservation{11,30,30}
	fmt.Printf("開始時刻：%d:%d、終了時刻：%d:%d", t.hour, t.min, t.hour + int((t.min+t.minutes)/60), (t.min + t.minutes)%60)
}
