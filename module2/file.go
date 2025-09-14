package main

import (
	"fmt"
	"os"
)

type method_of_files interface {
	uploadFile(string)
}

type fileStore struct {
	Filer *os.File
}

func (Fi fileStore) uploadFile(contant string) {
	fmt.Fprintf(Fi.Filer, contant)
}

func newFile(filename string) *fileStore {

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)

	if err != nil {
		panic(err)
	}

	return &fileStore{file}

}

func filer() {
	var log method_of_files
	log = newFile("riswan.py")
	log.uploadFile("print(\"hello world\")")

}
