package main

<<<<<<< HEAD
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
=======
import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
>>>>>>> 9514ac8a2c135a448a2b15a4b246dcd5d59ee7bf
