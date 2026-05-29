# Notes GO — M. CANETTI

## Idée générale

- séance surtout orientée TP
- notes pas censées être trop longues
- plutôt un squelette des notions importantes
- exemples surtout à récupérer dans les TP

## Notions vues / revues

- types de base : rappel
- array = tableau taille fixe
- slice = tableau dynamique
- `make` pour créer un slice
- `len` = longueur
- `cap` = capacité
- `append` pour ajouter
- copie de slice possible
- map = dictionnaire (comme python)

## Variables / constantes

- déclaration courte avec `:=`
- ex : `nom := "REDA"`
- double déclaration possible : `x, y := 10, 30`
- plusieurs types possibles si Go comprend
- constantes groupées avec `const (...)`
- `iota` pour incrémenter automatiquement

## Fonctions

- pas de surcharge en Go
- plusieurs valeurs de retour possibles
- souvent : résultat + erreur
- fonction variadique = nombre variable d'arguments
- closure = capture son environnement
- à ne pas confondre avec récursivité

## Boucles / conditions

- une seule boucle : `for`
- remplace `while`, `do while`, `foreach`
- `switch` à bien connaître
- `fallthrough` = forcer le cas suivant

## Struct / objet

- pas de classes comme en Java
- structs + méthodes
- pas d’héritage
- logique plus basée sur composition
- approche objet quand même

## Visibilité

- pas de `public/private/protected`
- visibilité selon la casse
- majuscule au début = exporté
- minuscule au début = non exporté
- camelCase / PascalCase à connaître

## Pointeurs

- pointe vers une adresse mémoire
- évite de copier des grosses structures
- utile pour modifier directement une donnée
- à revoir plus en détail plus tard

## defer

- exécution garantie à la fin d’une fonction
- même s’il y a une erreur
- utile pour fermer un fichier / libérer une ressource
- ordre LIFO : last in, first out

## Packages rappelés

- `fmt` : affichage, formatage
- `strings` : `Contains`, `ToUpper`, `Split`, `TrimSpace`
- `sort` : tri de slice / array
- éviter de recréer des fonctions déjà existantes

## À retenir pour les TP

- déclarations
- types
- `iota`
- fonctions
- erreurs
- closures
- `for`
- `switch`
- structs
- slices
- maps
- pointeurs
