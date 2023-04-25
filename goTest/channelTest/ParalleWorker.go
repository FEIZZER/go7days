package channelTest

import (
	"fmt"
	"time"
)

type Task struct {
	timeStamp int64
	paramsMap map[string]interface{}
	work      func()
}

var TaskChannel = make(chan *Task, 3)

func StartParalleWork() {
	for {
		task := <-TaskChannel
		process(task)
	}
}
func PushTaskChannel() {
	TaskChannel <- &Task{
		timeStamp: time.Now().Unix(),
		paramsMap: map[string]interface{}{},
		work: func() {
		},
	}
}
func process(task *Task) {
	fmt.Printf("start process task")
}
