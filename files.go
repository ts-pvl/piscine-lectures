package main

import (
	"fmt"
	"io"
	"os"
)

func files() {
	file, err := os.Open("files.go")
	if err != nil {
		fmt.Println("error")
		return
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println(string(bytes))

	err = file.Close()
	if err != nil {
		fmt.Println("error")
		return
	}
}
