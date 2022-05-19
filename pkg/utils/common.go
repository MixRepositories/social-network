package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func GetHashString(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func ParseStrToUint16(str string) (uint16, error) {
	var base = 16
	var size = 16
	value, err := strconv.ParseUint(str, base, size)
	println(value)
	if err != nil {
		return 0, err
	}

	return uint16(value), nil
}
