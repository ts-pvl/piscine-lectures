package main

import (
	"fmt"
	"io"
	"os"
)

func files() {
	file, err := os.Open("README.md")
	if err != nil {
		fmt.Println("error")
		return
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		file.Close()
		fmt.Println("error")
		return
	}

	fmt.Println(string(bytes))

	file.Close()
}
