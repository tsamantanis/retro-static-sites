package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

type Data struct {
	Content    string
	StylesPath string
}

func main() {
	filename := flag.String("file", "filename", "name of file to parse")
	directory := flag.String("dir", "directory", "name of directory of files to parse")
	flag.Parse()

	if *directory != "directory" {
		createManyHTML(*directory)
	}

	if *filename != "filename" {
		data := loadFileContent(*filename)
		createHTML("", strings.SplitN(*filename, ".", 2)[0], "template.tmpl", data)
	}

}

func loadFileContent(filename string) Data {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var dircount string = ""
	for i := 0; i < len(strings.SplitN(filename, "/", -1)); i++ {
		dircount = dircount + "../"
	}
	return Data{string(fileContents), dircount + "styles.css"}
}

func createHTML(dir, filename, templ string, data Data) {
	htmlFile, osErr := os.Create(dir + "/" + filename + ".html")
	if osErr != nil {
		log.Fatal(osErr)
	}

	t := template.Must(template.ParseFiles(templ))
	execErr := t.Execute(htmlFile, data)
	if execErr != nil {
		log.Fatal(execErr)
	}
}

func createManyHTML(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		path := directory + "/" + file.Name()
		stat, errStat := os.Stat(path)
		if errStat != nil {
			panic(err)
		}
		if !stat.IsDir() && strings.SplitN(stat.Name(), ".", 2)[1] == "txt" {
			data := loadFileContent(path)
			createHTML(directory, strings.SplitN(file.Name(), ".", 2)[0], "template.tmpl", data)
		} else if stat.IsDir() {
			createManyHTML(path)
		} else {
			fmt.Println("Error: File type does not match .txt")
		}
	}
}
