package handlers

import (
	"fmt"
)

type ReservationFinishHandler struct {

}

func (*h ReservationExtendHandler) Handle(params []string) {
	fmt.Printf(params)
}
