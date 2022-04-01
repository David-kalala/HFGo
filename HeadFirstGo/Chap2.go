package main

import (
	"fmt"
)

func main() {
	a := 1.0
	b := 3.0
	// Ici on utilise la fonction 'Printf' du package fmt pour afficher la reponse sans trop de chiffre apres la virgule
	fmt.Printf(" %0.2f", a/b)
	// Ici on fait la meme chose la seul diffrence que 'Sprintf' return un string au lieu de l'afficher
	// Donc on enregistre le return dans 'reponse' et on l'affiche
	reponse := fmt.Sprintf(" %0.2f", (a/b)+2)
	fmt.Println(reponse)
	fmt.Println()
	//On fait appel a la procedure helloGirl
	helloGirl()
	//On fait appel a la procedure add qui va additionner les deux nombres
	add(3, 10)

}

//Ici on creer une procedure qui affiche ' i love you girl ' et souligne la phrase
func helloGirl() {
	fmt.Println("Love you girl")
	for x := 0; x < 15; x++ {
		fmt.Print("-")
	}
	fmt.Println()
}

// Ici on declare une procedure 'add' qui additionne deux nombres passé en parametre.
func add(a int, b int) {
	result := a + b
	fmt.Println(a, "+", b, "=", result)
}
