package main

import "fmt"

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("Je m'appelle %s, j'ai %d ans et mon email est %s.", p.NomComplet(), p.Age, p.Email)
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"---- Fiche Employé ----\nNom : %s\nAge : %d\nEmail : %s\nAdresse : %s\nPoste : %s\nSalaire : %.2f €",
		e.NomComplet(),
		e.Age,
		e.Email,
		e.Format(),
		e.Poste,
		e.Salaire,
	)
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire = e.Salaire + (e.Salaire * pct / 100)
}

type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "TB"
	case e.Moyenne >= 14:
		return "B"
	case e.Moyenne >= 12:
		return "AB"
	default:
		return "P"
	}
}

func (e Etudiant) FicheEtudiant() string {
	return fmt.Sprintf(
		"---- Fiche Étudiant ----\nNom : %s\nAge : %d\nEmail : %s\nPromo : %s\nMoyenne : %.2f\nMention : %s",
		e.NomComplet(),
		e.Age,
		e.Email,
		e.Promo,
		e.Moyenne,
		e.MentionObtenue(),
	)
}

func main() {
	emp1 := Employe{
		Personne: Personne{
			Prenom: "Alice",
			Nom:    "Martin",
			Age:    30,
			Email:  "alice.martin@entreprise.com",
		},
		Adresse: Adresse{
			Rue:        "12 rue Victor Hugo",
			Ville:      "Lyon",
			CodePostal: "69000",
		},
		Poste:   "Développeuse",
		Salaire: 2800,
	}

	emp2 := Employe{
		Personne: Personne{
			Prenom: "Karim",
			Nom:    "Dupont",
			Age:    35,
			Email:  "karim.dupont@entreprise.com",
		},
		Adresse: Adresse{
			Rue:        "8 avenue Jean Jaurès",
			Ville:      "Villeurbanne",
			CodePostal: "69100",
		},
		Poste:   "Chef de projet",
		Salaire: 3500,
	}

	etu1 := Etudiant{
		Personne: Personne{
			Prenom: "Sofia",
			Nom:    "Benali",
			Age:    21,
			Email:  "sofia.benali@etu.fr",
		},
		Promo:   "B3 Informatique",
		Moyenne: 15.5,
	}

	etu2 := Etudiant{
		Personne: Personne{
			Prenom: "Yanis",
			Nom:    "Morel",
			Age:    22,
			Email:  "yanis.morel@etu.fr",
		},
		Promo:   "B3 Cybersécurité",
		Moyenne: 17.2,
	}

	emp1.AugmenterSalaire(10)

	fmt.Println(emp1.FicheEmploye())
	fmt.Println()
	fmt.Println(emp2.FicheEmploye())
	fmt.Println()
	fmt.Println(etu1.FicheEtudiant())
	fmt.Println()
	fmt.Println(etu2.FicheEtudiant())
}
