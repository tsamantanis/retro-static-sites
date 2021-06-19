package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Data struct {
	Content string
}

func main() {
	filename := flag.String("file", "filename", "name of file to parse")
	flag.Parse()
	data := loadFileContent(*filename)
	createHTML(strings.SplitN(*filename, ".", 2)[0], "template.tmpl", data)
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
