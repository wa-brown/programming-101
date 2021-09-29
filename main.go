package main

import (
	"fmt"
	"log"
	"time"

	"github.com/wa-brown/programming-101/concurrency"
)

func main() {
	// programming.VariableExample()

	//programming.LoopExample()

	//programming.FunctionExample()

	//warren := programming.StructExample()
	//programming.MethodExample(warren)

	// longProcess()
	pipelineExample()
}

func longProcess() {
	startTime := time.Now()

	instr := []string{}
	for i := 0; i < 5; i++ {
		instr = append(instr, fmt.Sprintf("%v", i))
	}

	outstr := concurrency.LongProcess(instr)
	_ = concurrency.LongerProcess(outstr)

	log.Printf("Time taken: %v", time.Since(startTime))
}

func pipelineExample() {
	startTime := time.Now()

	inchan := make(chan string, 100)
	go func() {
		defer close(inchan)
		for i := 0; i < 5; i++ {
			inchan <- fmt.Sprintf("%v", i)
		}
	}()

	longChan := concurrency.LongProcessPipeline(inchan)
	longerChan := concurrency.LongerProcessPipeline(longChan)

	// process the longerChan
	for str := range longerChan {
		log.Println(str)
	}

	log.Printf("Time taken: %v", time.Since(startTime))
}
