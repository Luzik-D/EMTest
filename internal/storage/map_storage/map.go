package storage

import ".github.com/Luzik-D/EMTest/internal/storage"

type Storage struct {
	st map[int]storage.Person
	id int
}

func New() *Storage {
	st := make(map[int]storage.Person)

	return &Storage{st, 1}
}

func (s *Storage) AddPerson(p storage.Person) {
	s.id++
	s.st[s.id] = p
}

func (s *Storage) GetPersons() []storage.Person {
	persons := make([]storage.Person, 0)

	for _, p := range s.st {
		persons = append(persons, p)
	}

	return persons
}
