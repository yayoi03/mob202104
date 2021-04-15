package handlers

import (
	"fmt"
)

type ReservationRegisterHandler struct {

}

func (*h ReservationRegisterHandler) Handle(params []string) {
	fmt.Printf(params)
}
