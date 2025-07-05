package datastore

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/murali-muniyan/golang/cache/cache"
)

type DataStore struct {
	data   map[int]string
	cache  cache.CacheI[int, string]
	logger log.Logger
}

func New() *DataStore {
	return &DataStore{
		data:   make(map[int]string),
		cache:  cache.New[int, string](10),
		logger: *log.New(os.Stdout, "", 0),
	}
}

func (d *DataStore) AddData(key int, val string) error {
	if _, exists := d.data[key]; exists {
		return errors.New("data already exists for the key")
	}

	d.data[key] = val

	if d.cache.IsFull() {
		return nil
	}

	err := d.cache.AddValWithTTL(key, val, time.Second*30)
	if err != nil {
		d.logger.Printf("error while adding data to cache: %v", err)
	}

	return nil
}

func (d *DataStore) GetVal(key int) (val string, err error) {
	val, err = d.cache.GetVal(key)
	if err == nil {
		log.Printf("got value from cache")
		return val, nil
	}

	log.Printf("error while getting data from cache: %v\n", err)

	val, ok := d.data[key]
	if !ok {
		return "", errors.New("data not exists in data store")
	}

	if d.cache.IsFull() {
		return val, nil
	}

	err = d.cache.AddValWithTTL(key, val, time.Second*30)
	if err != nil {
		log.Printf("error while adding to cache: %v\n", err)
	}

	fmt.Println(d.data)

	return val, nil
}

func (d *DataStore) Remove(key int) (val string, err error) {
	val, ok := d.data[key]
	if !ok {
		return "", errors.New("data not found in data store")
	}

	d.cache.Remove(key)

	fmt.Println(d.data)

	return val, nil
}

func (d *DataStore) Display() string {
	res := ""

	for k, v := range d.data {
		res += fmt.Sprintf("%d: %s\n", k, v)
	}

	fmt.Println(d.data)

	return res
}
