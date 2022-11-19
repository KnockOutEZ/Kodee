package utils

import "log"

func CheckErr(err error) {
	if err != nil {
		// notificationFunc("Error",err.Error())
		log.Fatal(err.Error())
	}
	return
}