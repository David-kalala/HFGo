package main

import (
	"fmt"
	"input"
)

func main() {
	//declare une variable myInt de type int
	var myInt int
	myInt = 23
	// declare un pointeur de type integer
	var myIntPointer *int
	//On stock l'adresse de myInt dans le pointeur myIntPointer
	myIntPointer = &myInt
	// declare un pointeur de pointer de type integer
	var pointerMyIntPointer **int
	//On stock l'adresse de myIntPointer dans le pointeur pointerMyIntPointer
	pointerMyIntPointer = &myIntPointer
	// declaration courte d'un pointeur de bool
	isPointer := true
	pointerIsPointer := &isPointer

	//On affiche l'address de myInt
	fmt.Println("myIntPointer ---> ", myIntPointer)
	fmt.Println("&myInt ---> ", &myInt)
	fmt.Println("---------------------------------------------")
	// on affiche l'address de myIntPointer
	fmt.Println("&myIntPointer ---> ", &myIntPointer)
	fmt.Println("pointerMyIntPointer --->", pointerMyIntPointer)
	fmt.Println("---------------------------------------------")
	//on affiche l'address de isPointer
	fmt.Println("&isPointer --->", &isPointer)
	fmt.Println("pointerIsPointer --->", pointerIsPointer)
	fmt.Println("***********************************************")
	//affiche le contenu du pointer
	fmt.Println("*myIntPointer --->", *myIntPointer)
	/////////////////////////////////////////// On ne peut regarder que le contenu d'un pointeur
	//fmt.Println("*myInt --->", *myInt)////// avec "*[name_pointer]"
	///////////////////////////////////////// Sinon ca fait une erreur
	// je modifie le contenu de la valeur myInt par le pointer myIntPointer
	*myIntPointer = 34
	fmt.Println(myInt)
	fmt.Println("*myIntPointer --->", *myIntPointer)
	//declare pointer de type bool
	var sisi *bool
	// initialisé et delacre de la variable yes
	yes := true
	//assigner l'address de yes dans le pointer sisi
	sisi = &yes
	//Appel la fonction poner avec comme parametre un pointeur
	poner(sisi)
	//Appel la fonction poner avec comme parametre une address
	poner(&yes)
	fmt.Println("***********************************************")

	//appel de la fonction double
	nombre := 3
	double(&nombre)
	fmt.Println("nombre --->", nombre)
	fmt.Println("***********************************************")
	//appel de la fonction negation sur un bool true
	ok := true
	negation(&ok)
	fmt.Println(ok)

	input.GetIn()

}

//Procedure poner qui prend en parametre un pointer de type bool
// return si le parametre est true ou false
func poner(oui *bool) {
	fmt.Println(*oui)
}

//Fonction qui return le double du nombre ecrit en parametre
func double(nbr *int) {
	*nbr *= 2
}

//Fonction qui fait la negation du boolean passé en parametre
func negation(a *bool) {
	*a = !*a
}
