package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetHashString(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
