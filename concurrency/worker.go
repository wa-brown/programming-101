package concurrency

import (
	"fmt"
	"time"
)

func Worker(id int, inchan <-chan string, outchan chan<- string) {
	for str := range inchan {
		fmt.Println("worker: ", id, " started job: ", str)
		time.Sleep(time.Second)
		fmt.Println("worker: ", id, " finished job: ", str)
		outchan <- str + " - done :)"
	}
}
