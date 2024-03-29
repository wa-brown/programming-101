package concurrency

import (
	"fmt"
	"log"
	"time"
)

// long process funcs
// example of a pipeline pattern
func LongProcess(instr []string) []string {
	outstr := []string{}
	for _, str := range instr {
		val := fmt.Sprintf("%v - %v", str, "longProcess")
		log.Println(val)
		outstr = append(outstr, val)
		time.Sleep(time.Second / 2)
	}

	return outstr
}

func LongerProcess(instr []string) []string {
	outstr := []string{}
	for _, str := range instr {
		val := fmt.Sprintf("%v - %v", str, "longerProcess")
		log.Println(val)
		outstr = append(outstr, val)
		time.Sleep(time.Second)
	}

	return outstr
}

// This is a template for all pipeline functions.
// func pipeline(in <-chan interface{}) <-chan interface{} {
// 	out := make(chan interface{}, 100)

// 	go func() {
// 		defer close(out)
// 		for val := range in {
// 			// Process the val and then add to the out channel
// 			out <- val
// 		}
// 	}()

// 	return out
// }

// example of a pipeline pattern
func LongProcessPipeline(inchan <-chan string) <-chan string {
	outchan := make(chan string, 100)

	go func() {
		defer close(outchan)
		for str := range inchan {
			val := fmt.Sprintf("%v - %v", str, "longProcessPipeline")
			log.Println(val)
			outchan <- val
			time.Sleep(time.Second / 2)
		}
	}()

	return outchan
}

func LongerProcessPipeline(inchan <-chan string) <-chan string {
	outchan := make(chan string, 100)

	go func() {
		defer close(outchan)
		for str := range inchan {
			val := fmt.Sprintf("%v - %v", str, "longerProcessPipeline")
			outchan <- val
			time.Sleep(time.Second)
		}
	}()

	return outchan
}
