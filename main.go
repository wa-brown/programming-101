package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/wa-brown/programming-101/concurrency"
)

func main() {
	// programming.VariableExample()

	//programming.LoopExample()

	//programming.FunctionExample()

	//warren := programming.StructExample()
	//programming.MethodExample(warren)

	// golang concurrency patterns

	// anonymous go func example
	// go func() {
	// 	fmt.Print("hello world")
	// }()

	// wait group example
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Print("hello world")
		wg.Done()
	}()

	wg.Wait() // Don't forget to add this otherwise it will not run the go func

	// concurrency patterns

	// longProcess()
	// pipelineExample()
	// fanOutAndIn()
	// poolParty()
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

func fanOutAndIn() {
	startTime := time.Now()

	inchan := make(chan string, 100)
	go func() {
		defer close(inchan)
		for i := 0; i < 5; i++ {
			inchan <- fmt.Sprintf("%v", i)
		}
	}()

	longChan1 := concurrency.LongProcessPipeline(inchan)
	longChan2 := concurrency.LongProcessPipeline(inchan)

	fanInChan1 := concurrency.MergeChannels(longChan1, longChan2)

	longerChan1 := concurrency.LongerProcessPipeline(fanInChan1)
	longerChan2 := concurrency.LongerProcessPipeline(fanInChan1)

	finalChan := concurrency.MergeChannels(longerChan1, longerChan2)

	// process the longerChan
	for str := range finalChan {
		log.Println(str)
	}

	log.Printf("Time taken: %v", time.Since(startTime))
}

func poolParty() {
	const jobNum = 5
	const workerNum = 20
	jobchan := make(chan string, jobNum)
	resultschan := make(chan string, jobNum)

	for id := 0; id < workerNum; id++ {
		go concurrency.Worker(id, jobchan, resultschan)
	}

	for i := 0; i < jobNum; i++ {
		jobchan <- fmt.Sprintf("i'm job: %v", i)
	}
	close(jobchan)

	for i := 0; i < jobNum; i++ {
		fmt.Println(<-resultschan)
	}
}
