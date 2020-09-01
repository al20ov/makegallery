package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type pathsData struct {
	Src string
	Alt string
}

type pageData struct {
	Title  string
	Images []pathsData
}

func getStaticImages(path string) (paths []pathsData) {
	files, _ := ioutil.ReadDir(path)
	num := len(files)
	paths = make([]pathsData, num)
	for i, v := range files {
		paths[i] = pathsData{path + v.Name(), v.Name()}
	}
	return
}

var (
	staticPath     *string = flag.String("s", "static/", "Path to the static assets")
	layoutTemplate *string = flag.String("t", "template/layout.html", "Path to the html template")
	outputPath     *string = flag.String("o", "index.html", "Output path")
	pageTitle      *string = flag.String("title", "Gallery", "Title of the gallery if applicable in the template")
)

func main() {
	flag.Parse()
	paths := getStaticImages(*staticPath)
	tmpl := template.Must(template.ParseFiles(*layoutTemplate))
	file, _ := os.Create(*outputPath)
	data := pageData{
		Title:  *pageTitle,
		Images: paths,
	}
	if len(data.Images) == 0 {
		fmt.Println("Warning: empty gallery")
	}
	tmpl.Execute(file, data)
	fmt.Println("Gallery successfully written to " + *outputPath)
}
