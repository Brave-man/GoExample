package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

var ErrNotExist = errors.New("file does not exist")

func IsNotExist(err error) bool {
	if pe, ok := err.(*os.PathError); ok  {
		err = pe.Err
	}
	return err == syscall.ENOENT || err == ErrNotExist
}

func main() {
	_, err := os.Open("/no/such/file")
	fmt.Println(os.IsNotExist(err))
	fmt.Println(IsNotExist(err))
}
