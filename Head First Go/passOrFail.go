package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	number, err := getFloat()
	if err != nil {
		log.Fatal()
	}
	passOrFail(&number)
	//appel de la procedure Hello du package greeting

}

//Une fonction qui retourne le float ecrit par l'user
func getFloat() (float64, error) {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	return number, nil
}

//Un fonction qui verifie si le contenu du pointeur en parametre est sup ou inf a 60
// >60 : pass  et <60 fail

func passOrFail(rep *float64) {
	var status string

	if *rep >= 60 {
		status = "you passed"
	} else {
		status = "you failed"
	}
	fmt.Println("A grade of", *rep, "means", status)
}
