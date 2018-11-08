package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

// ======== Global ======
var config Configuration

// ======== HTML function map ======
var funcmap = template.FuncMap{
	"unescaped": unescaped,
	"add":       add,
}

// =================================
func unescaped(x string) interface{} {
	return template.HTML(x)
}

func add(x, y int) int {
	return x + y
}

// =================================

// Convenience function for printing to stdout
func p(a ...interface{}) {
	fmt.Println(a)
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		p("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		p("Cannot get configuration from file", err)
	}
}

// parse HTML templates
// pass in a list of file names, and get a template
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	// templates := template.New("")
	// templates.Funcs(template.FuncMap{"unescaped": unescaped})
	templates := template.Must(template.New("").Funcs(funcmap).ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
