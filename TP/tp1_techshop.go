package main

import (
	"fmt"
	"strings"
)

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	Produits []Produit
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	for _, produit := range c.Produits {
		if produit.ID == p.ID {
			return fmt.Errorf("erreur : un produit avec l'ID %d existe déjà", p.ID)
		}
	}

	c.Produits = append(c.Produits, p)
	return nil
}

func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for _, p := range c.Produits {
		if p.ID == id {
			return p, nil
		}
	}

	return Produit{}, fmt.Errorf("erreur : aucun produit trouvé avec l'ID %d", id)
}

func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	var resultats []Produit

	for _, p := range c.Produits {
		if strings.EqualFold(p.Categorie, cat) {
			resultats = append(resultats, p)
		}
	}

	return resultats
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	nbModifies := 0

	for i, p := range c.Produits {
		if strings.EqualFold(p.Categorie, categorie) {
			nouveauPrix := p.Prix - (p.Prix * pct / 100)
			if nouveauPrix < 0 {
				nouveauPrix = 0
			}
			c.Produits[i].Prix = nouveauPrix
			nbModifies++
		}
	}

	return nbModifies
}

func (c *Catalogue) Vendre(id int, qte int) error {
	if qte <= 0 {
		return fmt.Errorf("erreur : la quantité doit être supérieure à 0")
	}

	for i, p := range c.Produits {
		if p.ID == id {
			if !p.Actif {
				return fmt.Errorf("erreur : le produit %s est inactif", p.Nom)
			}

			if p.Stock < qte {
				return fmt.Errorf("erreur : stock insuffisant pour %s (stock actuel : %d)", p.Nom, p.Stock)
			}

			c.Produits[i].Stock -= qte
			return nil
		}
	}

	return fmt.Errorf("erreur : aucun produit trouvé avec l'ID %d", id)
}

func (c Catalogue) Rapport() string {
	nbProduits := len(c.Produits)
	valeurTotale := 0.0
	nbActifs := 0

	for _, p := range c.Produits {
		valeurTotale += p.Prix * float64(p.Stock)
		if p.Actif {
			nbActifs++
		}
	}

	return fmt.Sprintf(
		"===== RAPPORT CATALOGUE =====\nNombre total de produits : %d\nProduits actifs : %d\nValeur totale du stock : %.2f €",
		nbProduits,
		nbActifs,
		valeurTotale,
	)
}

func afficherProduit(p Produit) {
	etat := "Inactif"
	if p.Actif {
		etat = "Actif"
	}

	fmt.Printf("ID: %d | %s %s | Catégorie: %s | Prix: %.2f € | Stock: %d | %s\n",
		p.ID, p.Marque, p.Nom, p.Categorie, p.Prix, p.Stock, etat)
}

func afficherListe(produits []Produit) {
	if len(produits) == 0 {
		fmt.Println("Aucun produit à afficher.")
		return
	}

	for _, p := range produits {
		afficherProduit(p)
	}
}

func main() {
	catalogue := Catalogue{}

	produitsInitiaux := []Produit{
		{ID: 1, Nom: "iPhone 15", Marque: "Apple", Prix: 969.00, Stock: 8, Categorie: "Smartphone", Actif: true},
		{ID: 2, Nom: "Galaxy S24", Marque: "Samsung", Prix: 899.00, Stock: 10, Categorie: "Smartphone", Actif: true},
		{ID: 3, Nom: "MacBook Air M3", Marque: "Apple", Prix: 1299.00, Stock: 5, Categorie: "Ordinateur", Actif: true},
		{ID: 4, Nom: "ThinkPad X1 Carbon", Marque: "Lenovo", Prix: 1499.00, Stock: 4, Categorie: "Ordinateur", Actif: true},
		{ID: 5, Nom: "AirPods Pro 2", Marque: "Apple", Prix: 279.00, Stock: 15, Categorie: "Audio", Actif: true},
	}

	for _, p := range produitsInitiaux {
		err := catalogue.AjouterProduit(p)
		if err != nil {
			fmt.Println(err)
		}
	}

	for {
		var choix int

		fmt.Println("\n===== TECHSHOP - MENU =====")
		fmt.Println("[1] Ajouter un produit")
		fmt.Println("[2] Chercher un produit")
		fmt.Println("[3] Appliquer des soldes")
		fmt.Println("[4] Vendre un produit")
		fmt.Println("[5] Voir le rapport")
		fmt.Println("[0] Quitter")
		fmt.Print("Votre choix : ")

		_, err := fmt.Scan(&choix)
		if err != nil {
			fmt.Println("Erreur de lecture du choix.")
			return
		}

		switch choix {
		case 1:
			var p Produit
			var actifInput string

			fmt.Print("ID : ")
			if _, err := fmt.Scan(&p.ID); err != nil {
				fmt.Println("Erreur de lecture de l'ID.")
				continue
			}

			fmt.Print("Nom : ")
			if _, err := fmt.Scan(&p.Nom); err != nil {
				fmt.Println("Erreur de lecture du nom.")
				continue
			}

			fmt.Print("Marque : ")
			if _, err := fmt.Scan(&p.Marque); err != nil {
				fmt.Println("Erreur de lecture de la marque.")
				continue
			}

			fmt.Print("Prix : ")
			if _, err := fmt.Scan(&p.Prix); err != nil {
				fmt.Println("Erreur de lecture du prix.")
				continue
			}

			fmt.Print("Stock : ")
			if _, err := fmt.Scan(&p.Stock); err != nil {
				fmt.Println("Erreur de lecture du stock.")
				continue
			}

			fmt.Print("Catégorie : ")
			if _, err := fmt.Scan(&p.Categorie); err != nil {
				fmt.Println("Erreur de lecture de la catégorie.")
				continue
			}

			fmt.Print("Actif (oui/non) : ")
			if _, err := fmt.Scan(&actifInput); err != nil {
				fmt.Println("Erreur de lecture du statut actif.")
				continue
			}

			p.Actif = strings.EqualFold(actifInput, "oui") || strings.EqualFold(actifInput, "true")

			err := catalogue.AjouterProduit(p)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Produit ajouté avec succès.")
			}

		case 2:
			var sousChoix int

			fmt.Println("[1] Chercher par ID")
			fmt.Println("[2] Chercher par catégorie")
			fmt.Print("Votre choix : ")

			if _, err := fmt.Scan(&sousChoix); err != nil {
				fmt.Println("Erreur de lecture.")
				continue
			}

			switch sousChoix {
			case 1:
				var id int
				fmt.Print("ID du produit : ")
				if _, err := fmt.Scan(&id); err != nil {
					fmt.Println("Erreur de lecture de l'ID.")
					continue
				}

				produit, err := catalogue.TrouverParID(id)
				if err != nil {
					fmt.Println(err)
				} else {
					afficherProduit(produit)
				}

			case 2:
				var categorie string
				fmt.Print("Catégorie : ")
				if _, err := fmt.Scan(&categorie); err != nil {
					fmt.Println("Erreur de lecture de la catégorie.")
					continue
				}

				resultats := catalogue.TrouverParCategorie(categorie)
				if len(resultats) == 0 {
					fmt.Println("Aucun produit trouvé dans cette catégorie.")
				} else {
					afficherListe(resultats)
				}

			default:
				fmt.Println("Sous-choix invalide.")
			}

		case 3:
			var categorie string
			var pct float64

			fmt.Print("Catégorie à solder : ")
			if _, err := fmt.Scan(&categorie); err != nil {
				fmt.Println("Erreur de lecture de la catégorie.")
				continue
			}

			fmt.Print("Pourcentage de réduction : ")
			if _, err := fmt.Scan(&pct); err != nil {
				fmt.Println("Erreur de lecture du pourcentage.")
				continue
			}

			if pct < 0 {
				fmt.Println("Erreur : le pourcentage doit être positif.")
				continue
			}

			nb := catalogue.AppliquerReduction(categorie, pct)
			fmt.Printf("%d produit(s) modifié(s).\n", nb)

		case 4:
			var id int
			var qte int

			fmt.Print("ID du produit à vendre : ")
			if _, err := fmt.Scan(&id); err != nil {
				fmt.Println("Erreur de lecture de l'ID.")
				continue
			}

			fmt.Print("Quantité à vendre : ")
			if _, err := fmt.Scan(&qte); err != nil {
				fmt.Println("Erreur de lecture de la quantité.")
				continue
			}

			err := catalogue.Vendre(id, qte)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Vente enregistrée avec succès.")
			}

		case 5:
			fmt.Println(catalogue.Rapport())

		case 0:
			fmt.Println("Fermeture de TechShop. À bientôt.")
			return

		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
