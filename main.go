package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type config struct {
	Level string
	Time  int
	Word  string
}

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/pictionary.html", servePictionary)

	log.Println(fmt.Sprintf("Listening on :%s...", port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {

	level := r.URL.Query().Get("level")
	time := r.URL.Query().Get("time")

	timeValue, err := strconv.Atoi(time)
	if err != nil && time != "" {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	config := config{
		Level: level,
		Time:  timeValue,
	}

	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", "index.html")
	ww := filepath.Join("templates", "widget-word.html")

	tmpl, err := template.ParseFiles(lp, fp, ww)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	getNextWord(&config)
	executeTemplate(w, tmpl, config)
}

func servePictionary(w http.ResponseWriter, r *http.Request) {

	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", r.URL.Path)

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	executeTemplate(w, tmpl, nil)

}

func executeTemplate(w http.ResponseWriter, t *template.Template, data interface{}) {
	err := t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func getNextWord(config *config) {
	if config.Level != "" && config.Time > 0 {
		word, err := getWord(config.Level)
		if err != nil {
			log.Println(err)
			config.Word = "Error"
			return
		}
		config.Word = word
		return
	}
	config.Word = ""
}
