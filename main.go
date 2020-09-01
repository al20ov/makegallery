package main

import (
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

func main() {
	paths := getStaticImages("./static/")
	tmpl := template.Must(template.ParseFiles("layout.html"))
	file, _ := os.Create("./index.html")
	data := pageData{
		Title:  "images",
		Images: paths,
	}
	tmpl.Execute(file, data)
}
