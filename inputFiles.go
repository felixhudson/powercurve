package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// GetAllFiles will return all tcx files in this dir
func GetAllFiles() []string {
	files, err := ioutil.ReadDir("./tcx")
	if err != nil {
		log.Fatal(err)

	}
	result := make([]string, 0)

	for _, file := range files {
		if strings.Index(file.Name(), "tcx") == len(file.Name())-3 {
			fmt.Println(file.Name())
			result = append(result, file.Name())
		}

	}
	return result

}
