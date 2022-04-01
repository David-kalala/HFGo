// J'AI FAIS UN PETIT JEU DE DEVINETTE

// Il  y a toujours un package main, ca veut dire que tout le reste du code appartiendra au package main
package main

import (
	// ici on importe des package qu'on va utiliser

	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {

	p, _ := template.ParseGlob("*.html")

	type theGame struct {
		Answer   int
		Attempts int
		Random   int
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var nombre int
		var essaie = 10

		x := time.Now().Unix()
		rand.Seed(x)
		nbAleatoire := rand.Intn(101) + 1

		data := theGame{nombre, essaie, nbAleatoire}

		p.ExecuteTemplate(w, "game", data)

	})
	Doscour, _ := os.Getwd()
	fileServer := http.FileServer(http.Dir(Doscour + "/Jeu"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	fmt.Print(Doscour)
	http.ListenAndServe(":8000", nil)
}
