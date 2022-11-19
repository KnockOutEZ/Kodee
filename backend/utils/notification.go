package utils

import "gopkg.in/toast.v1"

func NotificationFunc(title, message string) {
	notification := toast.Notification{
		AppID:   "Kodee",
		Title:   title,
		Message: message,
		// Audio: toast.Default,
		// Icon:                "frontend/src/assets/wails.ico",
		// Actions:             []toast.Action{{"protocol", "I'm a button", "https://www.google.com/search?q=qwe"}, {"protocol", "Me too!", ""}},
	}
	err := notification.Push()
	CheckErr(err)
	return
}