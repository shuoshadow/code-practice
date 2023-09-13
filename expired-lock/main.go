package main

import (
	lock "code-practice/expired-lock/lock"
	"fmt"
	"time"
)

func main() {
	l := lock.NewExpiredLock()
	//l.Lock(0)
	//
	//ch := make(chan struct{})
	//go func() {
	//	if err := l.Unlock(); err != nil {
	//		fmt.Println(err)
	//	}
	//	// 异常后关闭通道
	//	close(ch)
	//}()
	//
	//// 阻塞等待
	//<-ch

	l.Lock(1)
	<-time.After(1 * time.Second)
	l.Lock(0)
	if err := l.Unlock(); err != nil {
		fmt.Println(err)
	}
}
