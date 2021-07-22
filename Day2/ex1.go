package main

import (
	"fmt"
	"sync"
)

type Frequency struct { // Struct for overall frequency
	frequency map[string]int
	mutex sync.Mutex
}

func (frequency *Frequency) FrequencyCalculator(id int, jobs <-chan string, wg *sync.WaitGroup)  { // go routine for calculating frequency of a word

	defer wg.Done()
	word := <- jobs

	for _, letter := range word	{
			frequency.mutex.Lock()
			frequency.frequency[string(letter)] += 1
			frequency.mutex.Unlock()
	}
	fmt.Println(word," is done by worker ", id)

}


func main(){
	words := []string{"quick", "dog", "dog", "dog", "dog"}
	numWords := len(words)
	jobs := make(chan string)
	var finalfrequency = Frequency{make(map[string]int), sync.Mutex{}}
	var wg sync.WaitGroup

	fmt.Println("For given string ", words)

	for workernumber := 0; workernumber < numWords; workernumber++ {
		wg.Add(1)
		go finalfrequency.FrequencyCalculator(workernumber, jobs, &wg)
	}
	for position := 0; position < numWords; position++ {
		jobs <- words[position]
	}

	wg.Wait()

	fmt.Println("Frequency of each letter in all the strings are")
	fmt.Println(finalfrequency.frequency)

}