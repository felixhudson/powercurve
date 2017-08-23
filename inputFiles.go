package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// return all tcx files in this dir
func GetAllFiles() []string {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)

	}
	result := make([]string, 0)

	for _, file := range files {
		if strings.Index(file.Name(), "tcx") >= 0 {
			fmt.Println(file.Name())
			result = append(result, file.Name())
		}

	}
	return result

}