package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

type html struct {
	Name       string
	Occupation string
}

func main() {
	paths := []string{
		"/home/judsen/Code/Resume/config/config.html",
	}
	var data html
	var err error
	Judhtml := template.Must(template.New("config.html").ParseFiles(paths...))

	file, err := ioutil.ReadFile("/home/judsen/Code/Resume/config/Jud.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("%+v\n", data)

	out, err := os.Create("/home/judsen/Code/Resume/web/resume.css")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	err = Judhtml.Execute(out, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("success")

}
