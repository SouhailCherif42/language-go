package main

import (
	"fmt"
)

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("erreur : division par zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("erreur : operation inconnue")
	}
}

func creerOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a, b float64) float64 {
			return a + b
		}
	case "-":
		return func(a, b float64) float64 {
			return a - b
		}
	case "*":
		return func(a, b float64) float64 {
			return a * b
		}
	case "/":
		return func(a, b float64) float64 {
			return a / b
		}
	default:
		return func(a, b float64) float64 {
			return 0
		}
	}
}

func main() {
	var a, b float64
	var op string

	fmt.Println("Calculatrice Go")
	fmt.Println("Entrez : nombre nombre operation")
	fmt.Println("Exemple : 10 5 +")
	fmt.Println("Tapez Ctrl+C pour quitter")

	for {
		fmt.Print("> ")
		_, err := fmt.Scan(&a, &b, &op)
		if err != nil {
			fmt.Println("Erreur de lecture :", err)
			return
		}

		if op == "quit" {
			break
		}

		resultat, err := operer(a, b, op)
		if err != nil {
			fmt.Println(err)
			continue
		}

		operation := creerOperation(op)
		fmt.Printf("Résultat : %.2f\n", operation(a, b))
		fmt.Printf("Vérification operer() : %.2f\n", resultat)
	}
}
