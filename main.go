package main

import (
	"ex7/urlchecker"
	"fmt"
	"log"
	"time"
)

func main() {
	domains := []string{
		"walmart.com",
		"google.com",
		"amazon.com",
		"wikipedia.com",
		"youtube.com",
		"yahoo.com",
		"msn.com",
		"netflix.com",
		"crystal-lang.org",
	}

	checkTime(func() {
		urlchecker.WithMultipleChannels(domains)
	})

	checkTime(func() {
		fmt.Println("\n****** WithMultipleChannels ******")
		urlchecker.WithSingleChannel(domains)
	})

	checkTime(func() {
		fmt.Println("\n****** Sequentielly ******")
		urlchecker.Sequentielly(domains)
	})

	checkTime(func() {
		fmt.Println("\n****** WithWaitGroup ******")
		urlchecker.WithWaitGroup(domains)
	})
}

func checkTime(f func()) {
	startTime := time.Now().UnixNano()
	f()
	endTime := time.Now().UnixNano()
	duration := float64(endTime-startTime) / 1000000
	log.Printf("\n------ Duration : %v ms\n\n", duration)
}
