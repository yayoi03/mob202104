package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/yayoi03/mob202104/handlers"
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
	Handlers := map[string]handlers.Handler{
		"予約登録": handlers.ReservationRegisterHandler{},
		"利用開始": handlers.ReservationStartHandler{},
		"予約延長": handlers.ReservationExtendHandler{},
		"利用終了": handlers.ReservationFinishHandler{},
	}

	for {
		s := strings.Split(nextLine(), " ")
		if len(s) == 0 {
			return
		}
		
		eventName = s[0]
		if val, ok := dict[eventName]; ok {
			handler = Handlers[eventName]
			handler.Handle(s)
		}
	}

	// for _,t := range reservations {
	// 	fmt.Printf("開始時刻：%02d:%02d、終了時刻：%02d:%02d\n", t.hour, t.min, t.hour+int((t.min+t.minutes)/60), (t.min+t.minutes)%60)
	// }
}
