package main

import (
	"feiWebframework/goTest/channelTest"
	"time"
)

func main() {
	go channelTest.Test()
	channelTest.Test2()
	time.Sleep(10 * time.Second)
}
