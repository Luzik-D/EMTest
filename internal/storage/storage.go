package storage

type Person struct {
	FullName
	Age    string
	Sex    string
	Nation string
}

type FullName struct {
	Name       string
	Surname    string
	Patronimic string
}

type Storage struct {
	st map[int]Person
	id int
}

func New() *Storage {
	st := make(map[int]Person)

	return &Storage{st, 1}
}

func (s *Storage) AddPerson(p Person) {
	s.id++
	s.st[s.id] = p
}

func (s *Storage) GetPersons() []Person {
	persons := make([]Person, 0)

	for _, p := range s.st {
		persons = append(persons, p)
	}

	return persons
}
