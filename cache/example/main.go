package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/murali-muniyan/golang/cache/example/datastore"
)

func main() {
	d := datastore.New()

	demoDataStore(d)
}

func demoDataStore(d *datastore.DataStore) {
	message := "DataStore\nChoose operation:\n" +
		"1. GetAll Data\n" +
		"2. Add Data\n" +
		"3. Remove Data\n" +
		"4. Get Data\n" +
		"5. Exit"

	s := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(message)

		input, err := s.ReadString('\n')
		if err != nil {
			log.Printf("error reading intput: %v", err)

			continue
		}

		var (
			key int
			val string
		)

		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println(d.Display())
		case "2":
			key, err = readKey(s)
			if err != nil {
				continue
			}

			val, err = readVal(s)
			if err != nil {
				continue
			}

			err = d.AddData(key, val)
			if err == nil {
				log.Printf("data added successfully. key: %d, val %s\n", key, val)
			}
		case "3":
			key, err = readKey(s)
			if err != nil {
				continue
			}

			val, err = d.Remove(key)
			if err == nil {
				log.Printf("data removed successfully. key: %d, val %s\n", key, val)
			}
		case "4":
			key, err = readKey(s)
			if err != nil {
				continue
			}

			val, err = d.GetVal(key)
			if err == nil {
				log.Printf("data got successfully. key: %d, val %s\n", key, val)
			}
		case "5":
			return
		default:
			fmt.Println(input)
			fmt.Printf("%T\n", input)
			log.Printf("invalid input. please enter valid input.\n")
		}

		if err != nil {
			log.Printf("error while performing operation: %v\n", err)
		}
	}
}

func readKey(s *bufio.Reader) (int, error) {
	fmt.Println("Enter key:")
	keyStr, err := s.ReadString('\n')
	if err != nil {
		log.Printf("error reading intput: %v", err)

		return -1, err
	}

	keyStr = strings.TrimSpace(keyStr)

	key, err := strconv.Atoi(keyStr)
	if err != nil {
		log.Printf("invalid intput: %v, should be an int", string(keyStr))

		return -1, err
	}

	return key, nil
}

func readVal(s *bufio.Reader) (string, error) {
	fmt.Println("Enter value:")
	val, err := s.ReadString('\n')
	if err != nil {
		log.Printf("error reading intput: %v", err)

		return "", err
	}

	val = strings.TrimSpace(val)

	return val, nil
}
