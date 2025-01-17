package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(input string) string {
	md5Hash := md5.New()
	defer md5Hash.Reset()
	md5Hash.Write([]byte(input))
	return hex.EncodeToString(md5Hash.Sum(nil))
}
