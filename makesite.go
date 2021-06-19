package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type Data struct {
	Content string
}

func main() {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(fileContents))
	data := Data{string(fileContents)}

	htmlFile, osErr := os.Create("first-post.html")
	if osErr != nil {
		log.Fatal(osErr)
	}

	t := template.Must(template.ParseFiles("template.tmpl"))
	execErr := t.Execute(htmlFile, data)
	if execErr != nil {
		log.Fatal(execErr)
	}
}
