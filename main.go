package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type CoordinationService interface {
	ObtainLock() error
	ReleaseLock() error
}

// implements CoordinationService
type FileCoordinationService struct {
	Dir string
}

func (s *FileCoordinationService) ObtainLock() error {
	var err error
	_, err = os.Stat("/tmp/coordinate.lock")
	if os.IsNotExist(err) {
		fmt.Println("creating lock")
		err = ioutil.WriteFile("/tmp/coordinate.lock", []byte("lock"), os.ModePerm)
	} else if os.IsExist(err) {
		err = errors.New("lock already exists")
	}
	return err
}

func (s *FileCoordinationService) ReleaseLock() error {
	err := os.Remove("/tmp/coordinate.lock")
	if err == nil || os.IsNotExist(err) {
		return nil
	}
	return err
}

func main() {
	svc := &FileCoordinationService{}
	fmt.Println("Getting lock")
	err := svc.ObtainLock()
	fmt.Printf("%#v\n", err)
	err = svc.ReleaseLock()
	fmt.Printf("release locked: %#v\n", err)
	fmt.Println("Getting lock")
	err = svc.ObtainLock()
	fmt.Printf("%#v\n", err)
	err = svc.ReleaseLock()
	fmt.Printf("release locked: %#v\n", err)
}
