package main

import "fmt"

const (
	IMCMaigreur = 18.5
	IMCNormal   = 25.0
	IMCSurpoids = 30.0
	Nom         = "Souhail"
)

func main() {
	poids := 78.0
	taille := 1.82

	imc := poids / (taille * taille)

	fmt.Printf("Bonjour %s !\n", Nom)
	fmt.Printf("Votre poids : %.2f kg\n", poids)
	fmt.Printf("Votre taille : %.2f m\n", taille)
	fmt.Printf("Votre IMC est de : %.2f\n", imc)

	if imc < IMCMaigreur {
		fmt.Println("Catégorie : Maigreur")
	} else if imc < IMCNormal {
		fmt.Println("Catégorie : Normal")
	} else if imc < IMCSurpoids {
		fmt.Println("Catégorie : Surpoids")
	} else {
		fmt.Println("Catégorie : Obésité")
	}
}
