package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// func MD5(str string) string {
// 	data := []byte(str)
// 	has := md5.Sum(data)
// 	md5str := fmt.Sprintf("%x", has)
// 	return md5str
// }

// func GetMD5Encode(data string) string {

// }

func GetMD5HashWithWrite(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetMD5HashWithSum(text string) string {
	hasher := md5.New()
	return hex.EncodeToString(hasher.Sum([]byte(text)))
}
