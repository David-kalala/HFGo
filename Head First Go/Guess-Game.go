// J'AI FAIS UN PETIT JEU DE DEVINETTE

// Il  y a toujours un package main, ca veut dire que tout le reste du code appartiendra au package main
package main

import (
	// ici on importe des package qu'on va utiliser
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// ici on affiche d'abord un saut de ligne grace au " printLN ", puis le reste quoi haha
	fmt.Print("I've chosen a random number between 1 and 100.")
	fmt.Println(" Can you guess it ?")

	// ici on declare une variable de type integer (entier)
	var essaie int
	// On cree une un nombre aleatoire
	// "x := " c'est une autre maniere de declarer une variable mais seulement lorsqu'on sait ce que la variable va
	// contenir
	x := time.Now().Unix()
	rand.Seed(x)
	nbAleatoire := rand.Intn(101) + 1
	// INDIQUE LE NOMBRE D'ESSAIES
	fmt.Println("You have 10 attempts.")
	// ici on creer une boucle qui dit qu'on va faire 10 tours tant que ' essaie ' est > ou = 1. essaie vaut 10
	// a chaque tour on fait ' essai-- ' ce qui veut dire essai = essai - 1 donc 10-1
	for essaie = 10; essaie >= 1; essaie-- {
		fmt.Print("Make a guess: ")
		// On demande a l'utilisateur d'entrer un nombre, en creant une variable qui enregister ce que
		//l'user va ecrire
		reader := bufio.NewReader(os.Stdin)
		// On lit ce que l'utilisateur ecrira jusqu'a ce qu'il appui sur la touche ENTREE
		// le ' err ' c'est une variable qu'on a creer pour lorsqu'il y a une erreur
		input, err := reader.ReadString('\n')
		// On verifie ce que l'user ecrit, qu'il n'y a pas d'erreur
		if err != nil {
			//le log.fatal(err) exit le programme et renvoie l'erreur s'il y a une erreur
			log.Fatal(err)
		}
		// On efface tous les espace et les saut de ligne que l'user a pu ecrire
		input = strings.TrimSpace(input)
		//On convertit le string en integer car apres on va le comparer a un nombre
		nombre, err := strconv.Atoi(input)
		//On verifie qu'il n'y a pas d'erreurs
		if err != nil {
			log.Fatal(err)
		}
		//On cree un boolean
		isNegative := false
		//On cree une boucle qui est vrai tant que le nombre est <0 et >100
		// isNegative est vrai dans la boucle
		for i := 0; !(isNegative); i++ {
			if nombre < 0 || nombre > 100 {
				fmt.Println("I guessed a number between 1 and 100 but you chose", nombre, "don't you wanna guess ?")
				fmt.Print("Try again : ")
				//On redemande a l'user de l'ecrire
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}
				input = strings.TrimSpace(input)
				nombre, err = strconv.Atoi(input)
				if err != nil {
					log.Fatal(err)
				}
				//Si le nombre est [0,100], le boolean devient vrai donc faux dans la boucle et on peut sortir de la boucle
			} else if nombre > 0 || nombre < 100 {
				isNegative = true
			}

		}

		// On verifie si le nombre de l'user est inferieur au nombre aleatoire et on affiche qlq chose si oui
		if nombre < nbAleatoire {
			fmt.Println("Your guess was LOW")
			// On verifie si le nombre de l'user est egale au nombre aleatoire
		} else if nombre == nbAleatoire {
			fmt.Println("Good Job !The number I've chosen was", nbAleatoire)
			//Si l'user entre le bon nombre on arrete le programme
			break
			// On verifie si le nombre de l'user est superieur au nombre aleatoire
		} else if nombre > nbAleatoire {
			fmt.Println("Your guess is HIGH")
		}
		// On verifie si le noombre d'essaie est depassé.. mais j'arrive pas a bien le faire
		if essaie == 1 {
			fmt.Println("Sorry.. You lost. The number I've chosen was", nbAleatoire)
		}
		//On affiche le nombre d'essaies restant
		fmt.Println("You have ", essaie-1, " tries left")

	}

}
