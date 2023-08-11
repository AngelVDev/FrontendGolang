package main

import (
	"html/template"
)

func main() {
	tmp := template.ParseFiles("index.html")
		tmp.Execute()
}
