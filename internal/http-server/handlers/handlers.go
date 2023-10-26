package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	st ".github.com/Luzik-D/EMTest/internal/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type PersonRepository interface {
	GetPersons() []st.Person
	GetPersonById(id int) st.Person
	//GetFiltered()
	CreatePerson(person st.Person)
	UpdatePersonById(id int, person st.Person)
	DeleteById(id int)
}

func GetPersons(log *logrus.Logger, rep PersonRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		persons := rep.GetPersons()

		log.Info("Get all persons")
		log.Debug(persons)
	}
}
func GetPerson(log *logrus.Logger, rep PersonRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Error("Invalid id ", err)
			return
		}

		person := rep.GetPersonById(id)

		personJSON, err := json.Marshal(person)
		if err != nil {
			log.Errorf("Failed to get person with id %d", id)
			w.Write([]byte("Person not found"))

			return
		}

		log.Info("Get person ", id)

		w.Write([]byte(personJSON))

	}
}

func DeletePerson(log *logrus.Logger, rep PersonRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Error("Invalid id ", err)
			return
		}

		rep.DeleteById(id)

		log.Info("Person deleted, id: ", id)

		log.Warn("write response")
	}
}

func CreatePerson(log *logrus.Logger, rep PersonRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var person st.Person

		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			log.Error("Invalid JSON in create person request: ", err)
			return
		}

		rep.CreatePerson(person)

		log.Info("Person created")
		log.Debug("Person: ", person)

		log.Warn("write response")
	}
}

func UpdatePerson(log *logrus.Logger, rep PersonRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var person st.Person

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Error("Invalid id")
			return
		}

		err = json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			log.Error("Invalid JSON in update person request ", err)
			return
		}

		rep.UpdatePersonById(id, person)

		log.Info("Update person, id: ", id)

		log.Warn("write response")
	}
}
