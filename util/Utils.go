package util

import (
	"crypto/md5"
	"encoding/hex"
	"go_database/config"
)

// md5加密工具
func MD5(text string) (res string) {
	data := []byte(text)

	md5Ctx := md5.New()
	md5Ctx.Write(data)

	encyData := md5Ctx.Sum(nil)
	return hex.EncodeToString(encyData)
}

func StrIsBlank(str string) (res bool) {
	if str == "" || len(str) == 0 {
		return true
	}
	return false
}

func ToJsonStr(json interface{}) {

}

func GetContextConfig() (contextConfig *config.ContextConfig) {
	return config.NewContextConfig()
}
