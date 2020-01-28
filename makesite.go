package main

import (
	"io/ioutil"
)

func main() {

}

func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func renderTemplate()
