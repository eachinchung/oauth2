// @Title        random
// @Description  生成随机数的工具
// @Author       Eachin
// @Date         2021/4/14 5:49 下午

package utils

import (
	"math/rand"
	"time"
)

// RandString 生成随机字符串
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
