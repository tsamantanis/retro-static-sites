package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type Data struct {
	Content string
}

func main() {
	data := loadFileContent("first-post.txt")
	createHTML("first-post", "template.tmpl", data)
}

func loadFileContent(filename string) Data {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return Data{string(fileContents)}
}

func createHTML(filename, templ string, data Data) {
	htmlFile, osErr := os.Create(filename + ".html")
	if osErr != nil {
		log.Fatal(osErr)
	}

	t := template.Must(template.ParseFiles(templ))
	execErr := t.Execute(htmlFile, data)
	if execErr != nil {
		log.Fatal(execErr)
	}
}
