package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"example.com/blog/database"
)

// MD5 md5加密
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	result := hex.EncodeToString(h.Sum(nil))
	return result
}

// 判断redis中是否存在这个key
func IpExists(key string) bool {
	exists, err := database.Rdb.Exists(context.Background(), key).Result()
	if err != nil {
		fmt.Println("err", err)
		return false
	}
	if exists == 1 {
		return true
	}
	return false
}
