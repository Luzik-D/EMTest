package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	".github.com/Luzik-D/EMTest/internal/storage"
	"github.com/sirupsen/logrus"
)

type ageApiJSON struct {
	Age int `json:"age"`
}

type genderApiJSON struct {
	Gender string `json:"gender"`
}

type countryApiJSON struct {
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

const (
	ageApiURL     = "https://api.agify.io"
	genderApiURL  = "https://api.genderize.io"
	countryApiURL = "https://api.nationalize.io"
)

func APIHandler(st *storage.Storage, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		person, err := createPersonFromRequest(r.Body, log)
		if err != nil {
			log.Error("Failed to create person throught API: ", err)
			return
		}

		st.AddPerson(person)

		log.Info("Person created and saved in storage")
		log.Info("Persons: ", st.GetPersons())
	}
}

func createPersonFromRequest(reqBody io.ReadCloser, log *logrus.Logger) (storage.Person, error) {
	var person storage.Person

	err := json.NewDecoder(reqBody).Decode(&person)
	if err != nil {
		return storage.Person{}, err
	}

	reqMap := getOpenApiReqs(person.Name)

	//TODO: use goroutines with errgroup.WithContext
	//TODO: replace getAge, getGender, getNationality single func
	/* smth like
	   respMap := getAdditionalInfo(reqMap), respMap["age"] <- age etc */

	log.Warn("GOROUTINES ARE NOT IMPLEMENTED")
	log.Warn("Replace with single func isn't implemented")
	age, err := getAge(reqMap["age"])
	gender, err := getGender(reqMap["gender"])
	country, err := getNationality(reqMap["country"])

	log.Debugf("Get age: %d from request: %s", age, reqMap["age"])
	log.Debugf("Get country: %s from request: %s", gender, reqMap["gender"])
	log.Debugf("Get gender: %s from request: %s", country, reqMap["country"])

	person.Age = age
	person.CountryID = country
	person.Gender = gender

	return person, nil
}

func getAge(req string) (int, error) {
	resp, err := http.Get(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	tmp := ageApiJSON{}

	err = json.NewDecoder(resp.Body).Decode(&tmp)
	if err != nil {
		return 0, err
	}

	return tmp.Age, nil
}

func getGender(req string) (string, error) {
	resp, err := http.Get(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	tmp := genderApiJSON{}

	err = json.NewDecoder(resp.Body).Decode(&tmp)
	if err != nil {
		return "", err
	}

	return tmp.Gender, nil
}

func getNationality(req string) (string, error) {
	resp, err := http.Get(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	tmp := countryApiJSON{}

	err = json.NewDecoder(resp.Body).Decode(&tmp)
	if err != nil {
		return "", err
	}

	return tmp.Country[0].CountryID, nil
}

func getOpenApiReqs(name string) map[string]string {
	reqMap := make(map[string]string)

	reqMap["age"] = buildOpenApiRequest(name, ageApiURL)
	reqMap["gender"] = buildOpenApiRequest(name, genderApiURL)
	reqMap["country"] = buildOpenApiRequest(name, countryApiURL)

	return reqMap
}

func buildOpenApiRequest(name string, apiURL string) string {
	query := url.Values{}
	query.Add("name", name)

	return apiURL + "?" + query.Encode()
}
