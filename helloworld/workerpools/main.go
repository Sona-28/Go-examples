package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <- chan int, results chan <- int, wg *sync.WaitGroup){
	for j:= range jobs{
		fmt.Println("worker ", id, " started job ",j)
		// time.Sleep(time.Second)
		fmt.Println("worker ",id," finished job ",j)
		results <- j*2
	}
	wg.Done()
}

func main(){
	start := time.Now()
	var wg sync.WaitGroup
	const numJobs = 20
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for w:=1;w<=20;w++{
		wg.Add(1)
		go worker(w,jobs,results, &wg)
	}
	for j:=1; j<=numJobs; j++{
		jobs <- j
	}
	close(jobs)
	for a:=1; a<=numJobs; a++{
		<- results
	}
	wg.Wait();
	end := time.Since(start)
	fmt.Println(end)
}