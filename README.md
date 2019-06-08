# Rapport TP IA Algorithmes génétiques
> Mattéo AUGER (numéro étudiant : 20150982)

# Programme d'algorithme génétique 
## Préambule

### Compiler le programme
Le programme a été codé en `golang`, vous trouverez dans le dossier `bin` un fichier déjà compilé. Cependant, pour recompiler le programme vous povez utiliser la commande `make genetic_queens`.

### Exétuer le programme
L'exécution du programme se fait selon deux modes.

Le premier mode exécute l'algorithme avec les paramètres données : 
```bash
./bin/genetic_queens.x 
<taille de l échiquier> 
<taille de la population> 
<chances de mutation> 
<pourcentage des meilleurs individus à garder> 
<pourcentage des pires individus à garder> 
<pourcentage des enfants à garder>
```
Pour exécuter l'algorithme pour un échiquier de 50 cases :
```bash
./bin/genetic_queens.x 50 10000 0.05 0.15 0.05 0.8
```

Le second mode permet d'effectuer un benchmark de l'algorithme génétique de 10x10 à 70x70 cases.
Le résultat sera écrit dans le fichier "donnees.txt".
Pour exécuter le benchmark : `./bin/genetic_queens.x benchmark`

## Explication de l'algorithme
### Représentation du plateau

### Principe de l'algorithme

#### Sélection
#### Mutation
#### Croisement

## Résultats

### Résultat benchmark algorithme génétique vs recuit simulé

![Resultat algo génétique](img/result.png)

On observe ici que le temps d'exécution monte de plus en plus jusqu'à 20 secondes pour un échiquier de taille 50x50 pour l'algorithme génétique, alors que le recuit simulé trouve cette solution en à peine 35ms.
Cette marge de performances est due au fait que l'algorithme génétique est peu adapté au problème des n-reines à cause de la définition d'un gène et de l'opération de croisement entre deux individus. En effet, dans le problème des n reines un gène n'est pas exprimé uniquement par la position d'une reine ce dernier dépend directement de la position des autres reines car la fonction de fitness joue sur le nombre de reines en prises. L'algorithme nécessite beaucoup de générations et stagne facilement ce qui engendre un temps d'exécution relativement long. Le recuit simulé en revanche permet de sortir rapidement de ces "stagnations" afin de mieux cibler l'optimum.

# Résolution du problème

## Représentation du triplet
La représentation d'un échiquier est **la même que pour l'algorithme**.

Soit `T = (I, O, B)` : 
* `I`: ensemble des états initiaux
* `O`: ensemble des opérateurs : **opérateur de permutation**
* `B`: ensemble des états finaux

On applique la fonction `enPrise(plateau)` sur chaque solution afin d'orienter la recherche. L'objectif est d'atteindre 0.

### Pour un exemple de N = 4

`I = {[ 1, 2, 3, 4 ]}` I est égal à un plateau dont les reines sont disposées en diagonale.
`O = {Permutation()}` O est égal à une opération de permutation de deux reines.
`P = {[ 2, 4, 1, 3 ], [ 3, 4, 1, 2 ]}` P est égal aux solutions pour le N donné. Une solution est représentée par `Enprise(plateau) = 0`.

## Système de résolution du problème


## Algorithme A* 
