package main

import (
	"fmt"
	"time"
)

func emit(wordChannel chan string, doneChannel chan bool) {
	words := []string{"Filthy", "Recruiters", "Keep", "Shitting", "On", "Migrant", "Developers"}

	i := 0

	t := time.NewTimer(3 * time.Second)

	for {
		select {
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-doneChannel: // receive message to close the channel
			doneChannel <- true  // acknowledge by sending true in the channel
			return
		case <- t.C:
			return
		}
	}
}
func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)

	for word := range wordCh {
		fmt.Printf("%s ", word)
	}

	doneCh <- true
	<- doneCh
}
