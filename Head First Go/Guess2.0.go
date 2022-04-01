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

func acceuil(w http.ResponseWriter, r *http.Request) {

	var nombre int
	var essaie = 10

	x := time.Now().Unix()
	rand.Seed(x)
	nbAleatoire := rand.Intn(101) + 1

	data := theGame{Answer: nombre, Attempts: essaie, Random: nbAleatoire}
	p, _ := template.ParseFiles("game.html")
	p.ExecuteTemplate(w, "game", data)

}

type theGame struct {
	Answer   int
	Attempts int
	Random   int
}

func main() {
	http.HandleFunc("/guess", acceuil)
	http.ListenAndServe(":8000", nil)
}
