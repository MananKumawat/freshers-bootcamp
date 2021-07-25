package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)


type Teacher struct { // Struct for teacher
	TotalRating float64
}

type Student struct { // Struct for student
	id int
}

func (student Student) getrating(wg *sync.WaitGroup, result chan<- float64) { // Function for getting rating from a student
	defer wg.Done()
	randomtime := rand.Intn(50)
	time.Sleep(time.Millisecond * time.Duration(randomtime))
	result <- 0.01 * float64(rand.Intn(1000))
	//teacher.mutex.Lock()
	//teacher.TotalRating += 0.01 * float64(rand.Intn(1000))

	//teacher.mutex.Unlock()
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Changing seed according to time
	const NumberOfStudents = 200
	var teacher = Teacher{0.0}
	var Students []Student
	var result = make(chan float64, NumberOfStudents)
	var wg sync.WaitGroup

	for StudentId :=0; StudentId<NumberOfStudents; StudentId++ {
		Students = append(Students, Student{StudentId})
	}
	for _, student := range Students {
		wg.Add(1)
		go student.getrating(&wg, result)
	}
	wg.Wait()

	for i := 0; i < NumberOfStudents; i++ {
		teacher.TotalRating += <- result
	}

	teacher.TotalRating /= float64(NumberOfStudents)
	teacher.TotalRating = math.Round(teacher.TotalRating*100)/100
	fmt.Println("TotalRating of class teacher", teacher.TotalRating)
}