package demo

import (
	"time"
)

/*
HelloWorld 测试
*/
func HelloWorld() (res string) {
	res = time.Now().Local().String()
	return
}
