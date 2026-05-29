/*
EXERCICE NOTÉ 3 — Gestionnaire de notes avec slices et niveaux

Contexte : Vous devez créer un programme CLI en Go qui lit plusieurs notes,
les stocke dans un slice, calcule une moyenne, puis affiche une mention et un niveau.

1. Créez un groupe de constantes avec iota.
2. Déclarez un slice de float64 pour stocker plusieurs notes.
3. Créez une fonction moyenne(notes []float64) float64.
4. Créez une fonction mentionEtNiveau(moy float64) (string, int).
5. Utilisez un switch avec fallthrough.
6. Dans main(), affichez les notes, la moyenne, la mention et le niveau.

Rendu : fichier notes.go sur GitHub
*/
package main

import "fmt"

const (
	NiveauDebutant = iota
	NiveauIntermediaire
	NiveauAvance
	NiveauExpert
)

func moyenne(notes []float64) float64 {
	if len(notes) == 0 {
		return 0
	}

	var somme float64
	for _, note := range notes {
		somme += note
	}

	return somme / float64(len(notes))
}

func mentionEtNiveau(moy float64) (string, int) {
	switch {
	case moy < 10:
		return "Insuffisant", NiveauDebutant
	case moy < 12:
		return "Passable", NiveauIntermediaire
	case moy < 16:
		return "Bien", NiveauAvance
	default:
		return "Très bien", NiveauExpert
	}
}

func afficherValidations(niveau int) {
	fmt.Println("Validations obtenues :")

	switch niveau {
	case NiveauExpert:
		fmt.Println("- Niveau Expert")
		fallthrough
	case NiveauAvance:
		fmt.Println("- Niveau Avancé")
		fallthrough
	case NiveauIntermediaire:
		fmt.Println("- Niveau Intermédiaire")
		fallthrough
	case NiveauDebutant:
		fmt.Println("- Niveau Débutant")
	default:
		fmt.Println("- Aucun niveau")
	}
}

func nomNiveau(niveau int) string {
	switch niveau {
	case NiveauDebutant:
		return "Débutant"
	case NiveauIntermediaire:
		return "Intermédiaire"
	case NiveauAvance:
		return "Avancé"
	case NiveauExpert:
		return "Expert"
	default:
		return "Inconnu"
	}
}

func main() {
	notes := []float64{12.5, 15.0, 9.5, 18.0}

	fmt.Println("Liste des notes :")
	for i, note := range notes {
		fmt.Printf("Note %d : %.2f\n", i+1, note)
	}

	moy := moyenne(notes)
	mention, niveau := mentionEtNiveau(moy)

	fmt.Printf("\nMoyenne : %.2f\n", moy)
	fmt.Printf("Mention : %s\n", mention)
	fmt.Printf("Niveau atteint : %s\n", nomNiveau(niveau))

	afficherValidations(niveau)
}
