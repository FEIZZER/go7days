package channelTest

import "fmt"

var channel = make(chan int)
var channelWithBuffer = make(chan int, 10)
var nilChannel chan int = nil

func Test() {
	for true {
		select {
		case <-nilChannel:
			fmt.Printf("nil channel\n")
		default:
			fmt.Printf("always\n")
		}
	}
	nilChannel <- 1
}
func Test2() {
	nilChannel <- 1
}
