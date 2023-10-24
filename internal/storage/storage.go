package storage

type Person struct {
	FullName
	Age    int    `json:"age,omitempty"`
	Sex    string `json:"sex,omitempty"`
	Nation string `json:"nation,omitempty"`
}

type FullName struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronimic string `json:"patronimic,omitempty"`
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
