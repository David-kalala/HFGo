package main

import "fmt"

func main() {

	//Je declare un tableau contenant 5 cases de type string
	//ici il y a 6 cas mais seulement 5 initialisé, la derniere case sera nombres[6]=""
	nombres := [6]string{"David kalala kabamabi",
		"Ness Kalubi kabamabi",
		"Anaelle kalubi kabambi",
		"abiagelle kalubi kabambi",
		"Adonis Kalubi Kabambi"}

	// a l'aide d'une boucle j'affiche le contenu de chaque case
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}
	//ici j'affiche la case d'indice 3. sachant que la premiere indice est 0.
	fmt.Println(nombres[3])
	// je modifie la case d'indice 2 puis je l'affiche
	nombres[2] = "Nesse le bg "
	fmt.Println(nombres[2])
	//ici je declare un tableau de 2 case de type pointeur de string
	// ce tableau contient a l'indice l'adress de 4eme case du tableau nombres
	voyons := [2]*string{&nombres[4]}
	// j'affiche le contenu de l'adresse et l'adresse de la premiere case du tableau de pointeur de string
	fmt.Println(*voyons[0], voyons[0])
	fmt.Println(&nombres[4])
	//Je declare un slice de type string
	famille := []string{
		"Papa",
		"Maman",
		"Junior",
		"Diane",
		"Esme",
		"Robert",
	}
	// j'affiche la taille de la slice famille

	fmt.Println("Le nombre des membres de la famille :", len(famille))
	// je declare une sous slice qui comprend les deux premieres cases de la slice famille
	parents := famille[0:2]
	// je declare une sous slice qui comprend les cinqs dernieres cases de la slice famille
	enfants := famille[2:6]
	// je declare une sous slice qui comprend la troiseime case de la slice famille
	aine := famille[2:3]
	// je declare une sous slice qui comprend la cinquieme case de la slice famille
	cadet := famille[5:6]

	// J'affiche toutes les slices
	fmt.Println("Famille", famille, "\nParents", parents, "\nEnfants", enfants, "\nL'ainé(e)", aine, "\nLe cadet", cadet)

	//Un bebe vient de  naitre !
	famille = append(famille, "babyJo")
	cadet = famille[6:7]
	fmt.Println("La famille s'aggrandit !", famille, "\nNouveau cadet", cadet)
}
