package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return  hex.EncodeToString(tempStr)
}

// 大写
func MD5Encode(data string) string {
	return  strings.ToUpper(Md5Encode(data))
}

// 加密
func EncodePassword(password, salt string) string {
	return Md5Encode(password + salt)
}
// 解密
func DecodePassword(password, salt, sourcePassword string) bool {
	return Md5Encode(password + salt) == sourcePassword
}
