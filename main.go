package main

import "ex7/urlchecker"

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
	}

	urlchecker.WithChannels(domains)

}
