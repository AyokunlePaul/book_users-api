package utils

import (
	"regexp"
	"time"
)

const (
	emailPattern  = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	format        = "2006-01-02T15:04:05Z"
	mySQLDbFormat = "2006-01-02 15:04:05"
)

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(emailPattern)
	return emailRegex.MatchString(email)
}

func GetDBTime() string {
	return time.Now().UTC().Format(mySQLDbFormat)
}

func GetTime() string {
	return time.Now().UTC().Format(format)
}
