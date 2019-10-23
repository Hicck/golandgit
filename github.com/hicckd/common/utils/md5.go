package utils

import (
	"crypto/md5"
	"fmt"
)
//生成hash
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	md5 := fmt.Sprintf("%x", hash.Sum(nil))
	return ToUpper(md5)
}
// 生成32位MD5
//func MD5(text string) string{
//   ctx := md5.New()
//   ctx.Write([]byte(text))
//   return hex.EncodeToString(ctx.Sum(nil))
//}
