package storages

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/vuongle/grpc/helpers"
	"github.com/vuongle/grpc/pb"
)

var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
	Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex          // for handle concurrency
	data  map[string]*pb.Laptop // for store laptops, like a database
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock() // acquire lock to prevent race condition
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	// deep copy
	other, err := helpers.DeepCopy(laptop)
	if err != nil {
		return err
	}

	// save to the map
	store.data[other.Id] = other

	return nil
}

func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.Lock() // acquire lock to prevent race condition
	defer store.mutex.Unlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	// return a deep copy of found laptop
	return helpers.DeepCopy(laptop)
}

func (store *InMemoryLaptopStore) Search(
	ctx context.Context,
	filter *pb.Filter,
	found func(laptop *pb.Laptop) error) error {

	store.mutex.Lock() // acquire lock to prevent race condition
	defer store.mutex.Unlock()

	// 1.go through all the books
	for _, laptop := range store.data {

		// simulate slow network or heavy processing
		time.Sleep(time.Second)
		log.Printf("checking laptop id %s", laptop.Id)
		// check timeout or cancel error to prevent the server still save the data while client has caneled or timed out
		// if not check and handle error here, the code still runs and save the data even if client has caneled or timed out
		if ctx.Err() == context.DeadlineExceeded || ctx.Err() == context.Canceled {
			log.Println("context is timeout or canceled")
			return errors.New("context is timeout or canceled")
		}

		if helpers.IsQualified(filter, laptop) {
			// deep copy
			other, err := helpers.DeepCopy(laptop)
			if err != nil {
				return err
			}

			// 2.Instead of waiting to give all the books after going through all the books, the server instead would return books in a streaming fashion, as soon as it finds one
			// by calling the callback so that the callback streams data
			err = found(other)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
