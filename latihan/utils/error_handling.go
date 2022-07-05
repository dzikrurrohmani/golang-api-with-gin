package utils

import "log"

func RaiseError(err error) {
	if err != nil {
		log.Println(err)
	}
}