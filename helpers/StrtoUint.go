package helpers

import (
	"log"
	"strconv"
)

func STRtoUINT(id string) (uint, error) {
	user_PID_int, err := strconv.Atoi((id))
	if err != nil {
		log.Panic("error ", err)
	}
	user_PID := uint(user_PID_int)
	return user_PID, nil
}
