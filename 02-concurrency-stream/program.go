package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

// for incremental testing
func main_test() {
	wg := &sync.WaitGroup{}
	dataCh := make(chan int)
	doneCh := make(chan struct{})

	// testing "Source"
	wg.Add(1)
	go Source(wg, "data1.dat", dataCh)
	/*
		go func() {
			for no := range dataCh {
				fmt.Println(no)
			}
			close(doneCh)
		}()
	*/

	// testing "Splitter"
	/*
		go func() {
			evenCh, oddCh := Splitter(dataCh)
			splitterWg := &sync.WaitGroup{}
			splitterWg.Add(1)
			go func() {
				defer splitterWg.Done()
				for evenNo := range evenCh {
					fmt.Println("Even No :", evenNo)
				}
			}()
			splitterWg.Add(1)
			go func() {
				defer splitterWg.Done()
				for oddNo := range oddCh {
					fmt.Println("Odd No :", oddNo)
				}
			}()
			splitterWg.Wait()
			close(doneCh)
		}()
	*/

	wg.Wait()
	close(dataCh)
	<-doneCh
}

func main() {
	wg := &sync.WaitGroup{}
	dataCh := make(chan int)

	wg.Add(1)
	go Source(wg, "data1.dat", dataCh)

	wg.Add(1)
	go Source(wg, "data2.dat", dataCh)

	evenCh, oddCh := Splitter(dataCh)
	evenSumCh := Sum(evenCh)
	oddSumCh := Sum(oddCh)
	done := Merger(evenSumCh, oddSumCh, "result.txt")

	wg.Wait()
	close(dataCh)

	<-done
	fmt.Println("Done")
}

func Source(wg *sync.WaitGroup, fileName string, dataCh chan<- int) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if no, err := strconv.Atoi(line); err == nil {
			dataCh <- no
		}
	}
}

func Sum(noCh <-chan int) <-chan int {
	sumCh := make(chan int)
	go func() {
		// defer close(sumCh)
		var result int
		for no := range noCh {
			result += no
		}
		sumCh <- result
	}()
	return sumCh
}

func Splitter(dataCh <-chan int) (<-chan int, <-chan int) {
	evenCh := make(chan int)
	oddCh := make(chan int)
	go func() {
		defer close(evenCh)
		defer close(oddCh)
		for no := range dataCh {
			if no%2 == 0 {
				evenCh <- no
			} else {
				oddCh <- no
			}
		}
	}()
	return evenCh, oddCh
}

func Merger(evenSumCh <-chan int, oddSumCh <-chan int, fileName string) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		for range 2 {
			select {
			case evenSum := <-evenSumCh:
				fmt.Fprintf(file, "Even Total : %d\n", evenSum)
			case oddSum := <-oddSumCh:
				fmt.Fprintf(file, "Odd Total : %d\n", oddSum)
			}
		}
		close(doneCh)
	}()
	return doneCh
}
