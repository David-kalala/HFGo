// J'AI FAIS UN PETIT JEU DE DEVINETTE

// Il  y a toujours un package main, ca veut dire que tout le reste du code appartiendra au package main
package main

import (
	// ici on importe des package qu'on va utiliser

	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type theGame struct {
	Answer   int
	Attempts int
	Random   int
}

func main() {

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		var nombre int
		var essaie = 10

		x := time.Now().Unix()
		rand.Seed(x)
		nbAleatoire := rand.Intn(101) + 1

		data := theGame{Answer: nombre, Attempts: essaie, Random: nbAleatoire}
		p, _ := template.ParseGlob("*t.html")
		p.ExecuteTemplate(w, "game", data)

	})

	fileServer := http.FileServer(http.Dir("/game-test.html"))
	http.Handle("/guess", http.StripPrefix("/guess", fileServer))
	http.ListenAndServe(":8000", nil)
}
