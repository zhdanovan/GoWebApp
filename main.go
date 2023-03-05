package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Character struct {
	Name        string
	Element     string
	Rarity      int
	Description string
}

func main() {
	http.HandleFunc("/", characterHandler)
	http.ListenAndServe(":8080", nil)
}

func characterHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://api.genshin.dev/characters")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(body))
	var characters []Character
	if err := json.Unmarshal(body, &characters); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create an HTML template and execute it with the character data
	tmpl := template.Must(template.ParseFiles("character.html"))
	if err := tmpl.Execute(w, characters); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
