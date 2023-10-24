package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	".github.com/Luzik-D/EMTest/internal/storage"
	"github.com/sirupsen/logrus"
)

const (
	ageApiURL = "https://api.agify.io"
)

type InfoAdder interface {
	AddPerson(storage.Person)
}

// func getAge(ch chan string) {
// 	time.Sleep(time.Second)
// 	ch <- "22"
// }

// func getNation(ch chan string) {
// 	time.Sleep(time.Second)

// 	ch <- "russian"
// }

// func getSex(ch chan string) {
// 	time.Sleep(time.Second)

// 	ch <- "male"
// }

// func getFullName() storage.FullName {
// 	return storage.FullName{"petr", "petrov", "petrovich"}
// }

// func AddInfo(st *storage.Storage) {
// 	var age, n, s string

// 	ageChan := make(chan string)
// 	nChan := make(chan string)
// 	sChan := make(chan string)

// 	fullname := getFullName()
// 	go getAge(ageChan)
// 	go getNation(nChan)
// 	go getSex(sChan)

// 	for i := 0; i < 3; i++ {
// 		select {
// 		case g1 := <-ageChan:
// 			age = g1
// 		case g2 := <-nChan:
// 			n = g2
// 		case g3 := <-sChan:
// 			s = g3
// 		}
// 	}

// 	st.AddPerson(storage.Person{
// 		FullName: fullname,
// 		Age:      age,
// 		Nation:   n,
// 		Sex:      s,
// 	})

// }

func buildAgeRequest(name string) string {
	query := url.Values{}
	query.Add("name", name)

	return ageApiURL + "?" + query.Encode()
}

func getAge(req string) (int, error) {
	resp, err := http.Get(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	tmp := struct {
		Age int `json:"age"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&tmp)
	if err != nil {
		return 0, err
	}

	return tmp.Age, nil
}

func createPersonFromRequest(reqBody io.ReadCloser, log *logrus.Logger) (storage.Person, error) {
	var person storage.Person

	err := json.NewDecoder(reqBody).Decode(&person)
	if err != nil {
		return storage.Person{}, err
	}

	ageReq := buildAgeRequest(person.Name)
	age, err := getAge(ageReq)
	log.Debugf("Get age: %d from request: %s", age, ageReq)

	if err != nil {
		return storage.Person{}, err
	}

	person.Age = age

	return person, nil
}

func APIHandler(st *storage.Storage, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		person, err := createPersonFromRequest(r.Body, log)
		if err != nil {
			log.Error("Failed to create person throught API: ", err)
			return
		}

		log.Info("Create Person: ", person)
		log.Info("Persons: ", st.GetPersons())
	}
}

// func APIHandler(st *storage.Storage, log *logrus.Logger) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var person storage.Person

// 		err := json.NewDecoder(r.Body).Decode(&person)
// 		if err != nil {
// 			log.Error(err)
// 			return
// 		}

// 		log.Info("person body: ", person)
// 		ageReq := buildAgeRequest(person.FullName.Name)
// 		fmt.Println(ageReq)
// 		resp, err := http.Get(ageReq)
// 		if err != nil {
// 			log.Error(err)
// 		}

// 		t := struct {
// 			Age int `json:"age"`
// 		}{}
// 		err = json.NewDecoder(resp.Body).Decode(&t)
// 		if err != nil {
// 			log.Error(err)
// 		}
// 		fmt.Println(t.Age)
// 		log.Info("get age: ", t.Age)
// 		//nat req
// 		//sex req

// 		st.AddPerson(person)

// 		fmt.Println(st.GetPersons())
// 	}
// }
