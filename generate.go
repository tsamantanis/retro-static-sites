package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gomarkdown/markdown"
)

type Data struct {
	Content    template.HTML
	StylesPath string
}

type MdData struct {
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

	fileNum, size := create(*filename, *directory)

	boldGreen.Print("Success!")
	fmt.Print(" Generated ")
	boldWhite.Print(fileNum)
	fmt.Printf(" pages (%.1fkB total)\n", size)

}

func create(filename, directory string) (int16, float64) {
	var fileNum int16 = 0
	var size float64 = 0
	if directory != "directory" {
		fileNum, size = createManyHTML(directory)
	}

	if filename != "filename" {
		data := loadFileContent(filename)
		fileNum, size = createHTML("", strings.SplitN(filename, ".", 2)[0], "template.tmpl", data)
	}
	return fileNum, size
}

func loadFileContent(filename string) Data {
	fileContents, err := os.ReadFile(filename)
	if strings.SplitN(filename, ".", 2)[1] == "md" {
		md := []byte(fileContents)
		fileContents = markdown.ToHTML(md, nil, nil)
	}
	handleError(err)
	var dircount string = ""
	for i := 0; i < len(strings.SplitN(filename, "/", -1)); i++ {
		dircount = dircount + "../"
	}
	return Data{template.HTML(fileContents), dircount + "styles.css"}
}

func createHTML(dir, filename, templ string, data Data) (int16, float64) {
	var path string = ""
	if dir != "" {
		path = dir + "/"
	}
	htmlFile, osErr := os.Create(path + filename + ".html")
	handleError(osErr)
	t := template.Must(template.ParseFiles(templ))
	execErr := t.Execute(htmlFile, data)
	handleError(execErr)
	return 1, calculateSize(path + filename + ".html")
}

func createManyHTML(directory string) (int16, float64) {
	files, err := os.ReadDir(directory)
	handleError(err)
	var counter int16 = 0
	var size float64 = 0
	for _, file := range files {
		path := directory + "/" + file.Name()
		stat, errStat := os.Stat(path)
		handleError(errStat)
		if !stat.IsDir() && (strings.SplitN(stat.Name(), ".", 2)[1] == "txt" || strings.SplitN(stat.Name(), ".", 2)[1] == "md") {
			counter++
			data := loadFileContent(path)
			c, s := createHTML(directory, strings.SplitN(file.Name(), ".", 2)[0], "template.tmpl", data)
			counter += c
			size += s
		} else if stat.IsDir() {
			createManyHTML(path)
		} else {
			boldMagenta.Print("Warning:")
			fmt.Print(" File ")
			boldWhite.Print(stat.Name())
			fmt.Println(" does not match type .txt or .md")
		}
	}
	return counter, size
}

func handleError(err error) {
	if err != nil {
		boldRed.Print("Error:")
		fmt.Print(err)
		panic(err)
	}
}

func calculateSize(filename string) float64 {
	file, errOpen := os.Open(filename)
	handleError(errOpen)
	size, errSeek := file.Seek(0, 2)
	handleError(errSeek)
	return float64(size) / 1024.0
}
