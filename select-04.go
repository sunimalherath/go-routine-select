package main

import (
	"fmt"
	"time"
)

func emit(theChannel chan chan string, doneChannel chan bool) {
	wordChannel := make(chan string)
	theChannel <- wordChannel

	defer close(wordChannel)

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
	theCh := make(chan chan string)
	doneCh := make(chan bool)

	go emit(theCh, doneCh)

	wordCh := <- theCh

	for word := range wordCh {
		fmt.Printf("%s ", word)
	}

	doneCh <- true
	<- doneCh
}
