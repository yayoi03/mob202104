package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

type Reservation struct{
	hour,min,minutes int
}

func s2i(s string) int {
	i,err := strconv.Atoi(s)
	if err!=nil{
		panic(err)
	}
	return i
}

func main() {
	nstring := nextLine() // 予約の数
	n := s2i(nstring)
	reservations := make([]*Reservation,0,n)
	for i:=0; i<n; i++{
		s := strings.Split(nextLine(), " ")
		r := &Reservation{s2i(s[0]),s2i(s[1]),s2i(s[2])}
		reservations = append(reservations,r)
	}
	for _,t := range reservations {
		fmt.Printf("開始時刻：%02d:%02d、終了時刻：%02d:%02d\n", t.hour, t.min, t.hour+int((t.min+t.minutes)/60), (t.min+t.minutes)%60)
	}
}
