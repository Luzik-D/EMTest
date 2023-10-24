package api

import (
	"time"

	".github.com/Luzik-D/EMTest/internal/storage"
)

type InfoAdder interface {
	AddPerson(storage.Person)
}

func getAge(ch chan string) {
	time.Sleep(time.Second)
	ch <- "22"
}

func getNation(ch chan string) {
	time.Sleep(time.Second)

	ch <- "russian"
}

func getSex(ch chan string) {
	time.Sleep(time.Second)

	ch <- "male"
}

func getFullName() storage.FullName {
	return storage.FullName{"petr", "petrov", "petrovich"}
}
func AddInfo(st *storage.Storage) {
	var age, n, s string

	ageChan := make(chan string)
	nChan := make(chan string)
	sChan := make(chan string)

	fullname := getFullName()
	go getAge(ageChan)
	go getNation(nChan)
	go getSex(sChan)

	for i := 0; i < 3; i++ {
		select {
		case g1 := <-ageChan:
			age = g1
		case g2 := <-nChan:
			n = g2
		case g3 := <-sChan:
			s = g3
		}
	}

	st.AddPerson(storage.Person{
		FullName: fullname,
		Age:      age,
		Nation:   n,
		Sex:      s,
	})

}
