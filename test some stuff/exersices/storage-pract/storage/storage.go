package storage

import (
	"errors"
	"fmt"
)

type Employee struct {
	Id     int
	Name   string
	Age    string
	Salary int
}

type DumbStorage struct{}

func NewDumbStorage() *DumbStorage {
	return &DumbStorage{}
}

func (s *DumbStorage) Insert(e Employee) error {
	fmt.Printf("вставка пользователя с id: %d прошла успешно\n", e.Id)
	return nil
}

func (s *DumbStorage) Get(id int) (Employee, error) {
	e := Employee{
		Id: id,
	}

	return e, nil
}

func (s *DumbStorage) Delete(id int) error {
	fmt.Printf("удаление пользователя с id: %d прошла успешно\n", id)
	return nil
}

type Storage interface {
	Insert(e Employee) error
	Get(id int) (Employee, error)
	Delete(id int) error
}

type MemoryStorage struct {
	data map[int]Employee
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]Employee),
	}
}

func (s *MemoryStorage) Insert(e Employee) error {
	s.data[e.Id] = e

	return nil
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	e, exists := s.data[id]
	if !exists {
		return Employee{}, errors.New("employee with such id doesn't exists")
	}

	return e, nil
}

func (s *MemoryStorage) Delete(id int) error {
	delete(s.data, id)
	return nil
}
