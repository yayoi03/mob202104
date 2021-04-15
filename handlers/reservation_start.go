package handlers

import (
	"fmt"
)

type ReservationStartHandler struct {

}

func (*h ReservationStartHandler) Handle(params []string) {
	fmt.Printf(params)
}
