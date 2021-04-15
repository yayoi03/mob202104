package handlers

import (
	"fmt"
)

type ReservationExtendHandler struct {}

func (*h ReservationExtendHandler) Handle(params []string) {
	fmt.Printf(params)
}
