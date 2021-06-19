package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Data struct {
	Content    string
	StylesPath string
}

var boldGreen *color.Color = color.New(color.FgGreen, color.Bold)
var boldMagenta *color.Color = color.New(color.FgMagenta, color.Bold)
var boldRed *color.Color = color.New(color.FgRed, color.Bold)
var boldWhite *color.Color = color.New(color.FgWhite, color.Bold)

func main() {
	filename := flag.String("file", "filename", "name of file to parse")
	directory := flag.String("dir", "directory", "name of directory of files to parse")
	flag.Parse()

	var fileNum int16 = 0
	if *directory != "directory" {
		fileNum = createManyHTML(*directory)
	}

	if *filename != "filename" {
		data := loadFileContent(*filename)
		fileNum = createHTML("", strings.SplitN(*filename, ".", 2)[0], "template.tmpl", data)
	}

	boldGreen.Print("Success!")
	fmt.Print(" Generated ")
	boldWhite.Print(fileNum)
	fmt.Println(" pages")
}

func loadFileContent(filename string) Data {
	fileContents, err := os.ReadFile(filename)
	handleError(err)
	var dircount string = ""
	for i := 0; i < len(strings.SplitN(filename, "/", -1)); i++ {
		dircount = dircount + "../"
	}
	return Data{string(fileContents), dircount + "styles.css"}
}

func createHTML(dir, filename, templ string, data Data) int16 {
	htmlFile, osErr := os.Create(dir + "/" + filename + ".html")
	handleError(osErr)
	t := template.Must(template.ParseFiles(templ))
	execErr := t.Execute(htmlFile, data)
	handleError(execErr)
	return 1
}

func createManyHTML(directory string) int16 {
	files, err := os.ReadDir(directory)
	handleError(err)
	var counter int16 = 0
	for _, file := range files {
		path := directory + "/" + file.Name()
		stat, errStat := os.Stat(path)
		handleError(errStat)
		if !stat.IsDir() && strings.SplitN(stat.Name(), ".", 2)[1] == "txt" {
			counter++
			data := loadFileContent(path)
			createHTML(directory, strings.SplitN(file.Name(), ".", 2)[0], "template.tmpl", data)
		} else if stat.IsDir() {
			createManyHTML(path)
		} else {
			boldMagenta.Print("Warning:")
			fmt.Print(" File ")
			boldWhite.Print(stat.Name())
			fmt.Println(" does not match type .txt")
		}
	}
	return counter
}

func handleError(err error) {
	if err != nil {
		boldRed.Print("Error:")
		fmt.Print(err)
		panic(err)
	}
}
