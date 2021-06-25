package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// On ouvre le ficher de donnner afin de le lire
	file, err := os.Open("data.txt")
	// Verifie s'il y a une erreur à l'ouverture du fichier
	if err != nil {
		log.Fatal(err)
	}
	//Creer un scanner pour le fichier
	scanner := bufio.NewScanner(file)
	//On une ligne du fichier
	for scanner.Scan() {
		//On affiche la ligne
		fmt.Println(scanner.Text())
	}
	//" Clo de file for free ressources ", ca retourne seulement une erreur
	err = file.Close()
	// Check erros
	if err != nil {
		log.Fatal(err)
	}
	//S'il y a une erreur en scannant le fichier
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
