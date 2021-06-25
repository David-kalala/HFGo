package main

import "fmt"

func main() {
	//Ici on declare et initialise un maps/dico de string avec un string comme indice
	tabPeriodique := map[string]string{
		"Li": "Lithium",
		"Na": "Sodium",
		"K":  "Potasium",
		"O":  "Oxygène",
		"I":  "Iode",
	}
	//Ici on declare et initialise un maps/dico d'int avec un string comme indice
	meilleurButeur := map[string]int{
		"Christiano": 5,
		"Forsberg":   3,
		"Lukaku":     3,
	}
	// on affiche un element de chaque maps
	// On remarque que pour le dico meilleur Buteur "Messi" vaut 0 mais c'est la valeur par defaut et non
	// la vraie valeur
	fmt.Println(meilleurButeur["Lukaku"], meilleurButeur["Messi"])
	fmt.Println(tabPeriodique["Li"])

}

func contient(nom string, dico *map[string]int) bool {
	valeur, ok := *dico[nom]
	if !ok {
		fmt.Println("il n'y a pas de", valeur)
	} else {
		fmt.Println()
	}
	return ok
}
