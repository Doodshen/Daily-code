/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-29 18:50:59
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-29 19:35:33
 * @FilePath: \Daily code\day3_select\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"time"
)

// 多路复用
func demo1() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	select {
	case <-c1:
		fmt.Println("一号协程就绪")

	case <-c2:
		fmt.Println("二号协程就绪")
	}

}

// 超时控制
func demo2() {
	channel := make(chan string)

	select {
	case data := <-channel:
		fmt.Println("Received:", data)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout occurred")
	}
}

func main() {
	demo1()
	demo2()
}
