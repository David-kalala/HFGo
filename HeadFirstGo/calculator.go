package main

import (
	"errors"
	"fmt"
)

// Ici on creer une fonction qui prend deux floats en argument
// La fonction retourne un float et une erreur si le diviseur vaut 0
func divide(dividende float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("Can't divide by 0")

	}
	return dividende / divisor, nil
}

// Ici on creer une fonction qui prend deux floats en argument et un 'rune' un charactere
//cette fonction retourne un float
//cette fonction simule une calculatrice
func calc(nb1 float64, op rune, nb2 float64) float64 {
	var result float64
	if op == '+' {
		result = nb1 + nb2
	} else if op == '-' {
		result = nb1 - nb2
	} else if op == '*' {
		result = nb1 * nb2
	} else if op == '/' {
		result = nb1 / nb2
	}
	return result
}
func main() {
	//On fait appel a la fonction cal et on stok le resultat dans la variable a
	a := calc(2.8, '/', 5)
	// On affiche a un formatant le resultat afin qu'il n y ait pas trop de chiffre apres la virgule
	fmt.Printf("%.2f", a)

	//On fait appel a la fonction divide et on stok le resultat dans la variable b
	b, err := divide(2.4, 0.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%.2f", b)
}
