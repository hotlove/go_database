package util

import (
	"crypto/md5"
	"encoding/hex"
)

// md5加密工具
func MD5(text string) (res string)  {
	data := []byte(text)

	md5Ctx := md5.New()
	md5Ctx.Write(data)

	encyData := md5Ctx.Sum(nil)
	return hex.EncodeToString(encyData)
}
