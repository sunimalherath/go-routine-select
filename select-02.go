package main

import "fmt"

func emit(wordChannel chan string, doneChannel chan bool) {
	words := []string{"Filthy", "Recruiters", "Keep", "Shitting", "On", "Migrant", "Developers"}

	i := 0

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
		}
	}
}
func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)

	for i := 0; i < 100; i++ {
		fmt.Printf("%s ", <-wordCh)
	}

	doneCh <- true
	<- doneCh
}
