/*
EXERCICE NOTÉ 3 — Gestionnaire de scores avec array, slice et niveaux

Contexte :
Vous devez créer un programme CLI en Go qui manipule un tableau fixe de scores,
crée un slice à partir de ce tableau, puis un second slice avec make,
avant d’afficher la moyenne et le niveau de réussite.

L’objectif est de démontrer les notions suivantes :
- array
- slice
- len()
- cap()
- boucle for
- switch avec fallthrough
- constantes avec iota

Consignes :
 1. Créez un groupe de constantes avec iota pour représenter les niveaux :
    Niveau1, Niveau2, Niveau3, Niveau4.

 2. Déclarez un array de 5 entiers contenant des scores,
    par exemple : [5]int{10, 14, 8, 17, 12}.

 3. Créez un slice à partir de cet array,
    par exemple : scoresSlice := scores[1:4],
    puis affichez ses valeurs avec une boucle for.

4. Affichez len() et cap() de l’array et du slice.

 5. Créez un second slice avec make,
    par exemple : bonus := make([]int, 3, 5),
    puis remplissez-le avec une boucle for.

 6. Créez une fonction moyenne(notes []int) float64
    qui calcule la moyenne d’un slice avec for.

 7. Créez une fonction niveau(moy float64) int
    qui retourne un niveau avec un switch.

 8. Créez une fonction afficherNiveaux(n int)
    qui utilise fallthrough pour afficher tous les niveaux validés
    jusqu’au niveau obtenu.

9. Dans main(), affichez :
  - l’array
  - le slice extrait
  - le slice créé avec make
  - len() et cap()
  - la moyenne
  - le niveau atteint

Rendu :
fichier scores.go sur GitHub
*/
package main

import "fmt"

const (
	Niveau1 = iota
	Niveau2
	Niveau3
	Niveau4
)

func moyenne(notes []int) float64 {
	if len(notes) == 0 {
		return 0
	}

	somme := 0
	for _, note := range notes {
		somme += note
	}

	return float64(somme) / float64(len(notes))
}

func niveau(moy float64) int {
	switch {
	case moy < 10:
		return Niveau1
	case moy < 12:
		return Niveau2
	case moy < 15:
		return Niveau3
	default:
		return Niveau4
	}
}

func nomNiveau(n int) string {
	switch n {
	case Niveau1:
		return "Niveau 1"
	case Niveau2:
		return "Niveau 2"
	case Niveau3:
		return "Niveau 3"
	case Niveau4:
		return "Niveau 4"
	default:
		return "Inconnu"
	}
}

func afficherNiveaux(n int) {
	fmt.Println("Niveaux validés :")

	switch n {
	case Niveau4:
		fmt.Println("- Niveau 4")
		fallthrough
	case Niveau3:
		fmt.Println("- Niveau 3")
		fallthrough
	case Niveau2:
		fmt.Println("- Niveau 2")
		fallthrough
	case Niveau1:
		fmt.Println("- Niveau 1")
	default:
		fmt.Println("- Fin de l'affichage")
	}
}

func main() {
	scores := [5]int{10, 14, 8, 17, 12}
	scoresSlice := scores[1:4]

	bonus := make([]int, 3, 5)
	for i := 0; i < len(bonus); i++ {
		bonus[i] = (i + 1) * 2
	}

	fmt.Println("Array scores :", scores)
	fmt.Println("Slice extrait :", scoresSlice)
	fmt.Println("Slice bonus avec make :", bonus)

	fmt.Printf("len(scores) = %d | cap(scores) = %d\n", len(scores), cap(scores))
	fmt.Printf("len(scoresSlice) = %d | cap(scoresSlice) = %d\n", len(scoresSlice), cap(scoresSlice))
	fmt.Printf("len(bonus) = %d | cap(bonus) = %d\n", len(bonus), cap(bonus))

	moy := moyenne(scoresSlice)
	n := niveau(moy)

	fmt.Printf("Moyenne du slice extrait : %.2f\n", moy)
	fmt.Printf("Niveau atteint : %s\n", nomNiveau(n))

	afficherNiveaux(n)
}
