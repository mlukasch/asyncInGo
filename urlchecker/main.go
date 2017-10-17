package urlchecker

import (
	"fmt"
	"net/http"
	"sync"
)

func checkDomain(domain string) bool {
	resp, err := http.Get(fmt.Sprintf("http://%s", domain))
	if err != nil {
		return false
	}
	return (resp.StatusCode == 200)
}

func WithWaitGroup(domains []string) {
	fmt.Println("****** WithWaitGroup ******")
	var w sync.WaitGroup
	w.Add(len(domains))

	for _, d := range domains {
		go func(domain string, wg *sync.WaitGroup) {
			ok := checkDomain(domain)
			fmt.Printf("* %s : %v\n", domain, ok)
			w.Done()
		}(d, &w)
	}
	w.Wait()
}

func Sequentielly(domains []string) {
	fmt.Println("****** Sequentielly ******")
	for _, d := range domains {
		ok := checkDomain(d)
		fmt.Printf("* %s : %v\n", d, ok)
	}
}
func WithChannels(domains []string) {
	fmt.Println("****** WithChannels ******")
	var channelList []chan string
	for _, d := range domains {
		ch := make(chan string)
		go func(domain string, channel chan string) {
			ok := checkDomain(domain)
			channel <- fmt.Sprintf("* %s : %v\n", domain, ok)
		}(d, ch)
		channelList = append(channelList, ch)
	}
	for _, channel := range channelList {
		fmt.Print(<-channel)
	}
}
