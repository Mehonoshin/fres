package utils

import (
	"fmt"
	"os"
	"io/ioutil"
)

func GoToDir(path string) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	file.Chmod(0755)
	file.Close()
}

func CreateDir(path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		// Check if directory exists
		fmt.Println("Can't create dir")
	}
}

func WriteToFile(path, content string) error {
	err := ioutil.WriteFile(path, []byte(content), 0755)
	return err
}
