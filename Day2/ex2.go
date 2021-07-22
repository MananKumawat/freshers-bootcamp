package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)


type Rating struct { // Struct for rating
	TotalRating float64
	mutex sync.Mutex
}

type Student struct { // Struct for student
	id int
}

func (student Student) getrating(wg *sync.WaitGroup, totalrating *Rating) { // Function for getting rating from a student
	defer wg.Done()
	time.Sleep(time.Millisecond * 50)
	totalrating.mutex.Lock()
	totalrating.TotalRating += 0.01 * float64(rand.Intn(1000))
	totalrating.mutex.Unlock()
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Changing seed according to time
	const NumberOfStudents = 200
	var totalrating = Rating{0.0, sync.Mutex{}}
	Students := []Student{}
	var wg sync.WaitGroup

	for StudentId :=0; StudentId<NumberOfStudents; StudentId++ {
		Students = append(Students, Student{StudentId})
	}
	for _, student := range Students {
		wg.Add(1)
		go student.getrating(&wg, &totalrating)
	}
	wg.Wait()
	totalrating.TotalRating /= float64(NumberOfStudents)
	totalrating.TotalRating = math.Round(totalrating.TotalRating*100)/100
	fmt.Println(totalrating.TotalRating)
}