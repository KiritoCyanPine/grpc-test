package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	grpctest "github.com/kiritocyanpine/grpctest/pb/proto"
)

var ErrAlreadyExist = errors.New("key already exists in store")
var ErrNotFound = errors.New("id not found in store")

type LaptopStore interface {
	Save(laptop *grpctest.Laptop) error

	Find(id string) (*grpctest.Laptop, error)

	GetAll() []*grpctest.Laptop
}

type InMemoryLaptopStore struct {
	mutex *sync.RWMutex
	data  map[string]*grpctest.Laptop
}

func (store *InMemoryLaptopStore) Save(laptop *grpctest.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	_, exist := store.data[laptop.Id]
	if exist {
		return ErrAlreadyExist
	}

	// deep copy
	cache := &grpctest.Laptop{}

	err := copier.Copy(cache, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy data: %s", err)
	}

	store.data[laptop.Id] = cache

	return nil
}

func (store *InMemoryLaptopStore) Find(id string) (*grpctest.Laptop, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	val, exist := store.data[id]
	if !exist {
		return nil, ErrNotFound
	}

	cache := &grpctest.Laptop{}

	err := copier.Copy(cache, val)
	if err != nil {
		return nil, fmt.Errorf("cannot copy data: %s", err)
	}

	return cache, nil
}

func (store *InMemoryLaptopStore) GetAll() []*grpctest.Laptop {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	s := []*grpctest.Laptop{}
	for _, v := range store.data {
		s = append(s, v)
	}

	return s
}

func CreateInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		&sync.RWMutex{},
		make(map[string]*grpctest.Laptop),
	}
}
