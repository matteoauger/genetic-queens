# Rapport TP IA Algorithmes génétiques
> Mattéo AUGER (numéro étudiant : 20150982)

## Préambule

### Compiler le programme
Le programme a été codé en `golang`, vous trouverez dans le dossier `bin` un fichier déjà compilé. Cependant, pour recompiler le programme vous povez utiliser la commande `make genetic_queens`.
Pour compiler le recuit simulé : `make recuit_simule`.

### Recuit simulé : Exécuter le programme
`./bin/recuit-simule.x`

### Algorithme génétique : Exétuer le programme
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

## Explication de l'algorithme génétique
### Représentation du plateau
Le plateau est représenté dans un tableau à 1 dimension de taille N.
Ce dernier contient la colonne de la reine pour chaque ligne.
Par exemple, `[2, 1, 0, 3]` représente le tableau suivant : 
```
 |   |   | X |   | 
 |   | X |   |   | 
 | X |   |   |   | 
 |   |   |   | X |
```

### Principe de l'algorithme
L'algorithme génétique est un algorithme d'optimisation méta-heuristique basé sur la **génétique** et le mécanisme de **sélection naturelle**. Une solution devient un **individu**, contenu dans une population. Chaque population est modifiée au cours des différentes **générations**. Chaque individu possède des **gènes** qui pourront être sujet à **mutation**. Deux individus peuvent se **reproduire** et la population totale sera sujette à une sélection basée sur le principe de *survival of the fittest* (survie du plus apte). 

Une fonction de **fitness** est définie afin de juger l'aptitude de chaque individu. Il s'agit en quelque sorte d'une "note" que l'on attribute à une solution. 
La sélection gardera un certain pourcentage des meilleurs mais également des pires afin de garder une certaine **diversité** ce qui permet de ne pas s'échouer dans des extremums locaux.


#### Mutation
La mutation s'effectue de la manière suivante : 
Après reproduction, un enfant peut voir un de ses gênes mutés. La mutation dans le cas du problème des n reines se traduit par une permutation aléatoire dans le plateau de la solution.

#### Croisement
La sélection des individus pour le croisement se fait à l'aide de la roue de la fortune biaisée se basant sur la fitness de chaque individu. Cela permet de sélectionner les meilleurs individus mais aussi de laisser une chance aux moins bon de participer au processus de reproduction.
Le mécanisme de croisement permet la reproduction entre deux individus.
Afin d'effectuer cette reproduction, nous basons le plateau de l'enfant sur le plateau du premier parent.
Nous cherchons ensuite une valeur aléatoire dans le plateau du second parent, et nous permutons la position de cette valeur avec la position de cette même valeur dans le tableau enfant.

```go
func (pair QueenSolutionPair) Crossover() QueenSolution {
	n := len(pair.A.Board)

	childBoard := make([]int, n)

	copy(childBoard, pair.A.Board)

	randIndex := rand.Intn(len(pair.B.Board))
	childIndex := util.Find(childBoard, pair.B.Board[randIndex])

	childBoard[randIndex], childBoard[childIndex] = childBoard[childIndex], childBoard[randIndex]

	return CreateSolution(childBoard)
}
```
Exemple : fils de `[1,2,3,4]` et `[2,4,1,3]` (permutation avec `4`) : `[1,4,3,2]`

#### Sélection 
La sélection s'effectue de manière simple : On garde un certain pourcentage des meilleurs, des pires et des enfants de la génération actuelle.

## Résultats

### Résultat benchmark algorithme génétique vs recuit simulé

![Resultat algo génétique](img/result.png)

On observe ici que le temps d'exécution monte de plus en plus jusqu'à 9,3 secondes pour un échiquier de taille 50x50 pour l'algorithme génétique, alors que le recuit simulé trouve cette solution en à peine 35ms.
Cette marge de performances est due au fait que l'algorithme génétique est peu adapté au problème des n-reines à cause de la définition d'un gène et de l'opération de croisement entre deux individus. En effet, dans le problème des n reines un gène n'est pas exprimé uniquement par la position d'une reine ce dernier dépend directement de la position des autres reines car la fonction de fitness joue sur le nombre de reines en prises. L'algorithme nécessite beaucoup de calculs pour chaque génération alors que le recuit simulé ne nécessite qu'une permutation à chaque changement de température. 
Nous pouvons donc facilement déduire que le recuit simulé est beaucoup plus adapté au problème des n-reines que l'algorithme génétique. Pour une définition de gènes et des opérations de croisements plus simples, il est tout à fait possible que l'algorithme génétique soit meilleur que le recuit simulé étant donné la vitesse de convergence vers une solution. 


# Résolution du problème avec A*

## Représentation du triplet
La représentation d'un échiquier est **la même que pour l'algorithme**.

Soit `T = (I, O, B)` : 
* `I`: ensemble des états initiaux
* `O`: ensemble des opérateurs : **opérateur de permutation**
* `B`: ensemble des états finaux

On applique la fonction `enPrise(plateau)` sur chaque solution afin d'orienter la recherche. L'objectif est d'atteindre 0.

### Pour un exemple de N = 4

`I = {[ 1, 2, 3, 4 ]}` I est égal à un plateau dont les reines sont disposées en diagonale.
`O = {Permutation(reine1, reine2)}` O est égal à une opération de permutation de deux reines.
`P = {[ 2, 4, 1, 3 ]}` P est égal aux solutions pour le N donné. Une solution est représentée par `Enprise(plateau) = 0`.

## Heuristique 

* `h(noeud) = cout + distance entre le noeud et la fin`
* `h(noeud) = noeud.cout + noeud.enPrise`
On ajoute le coût du noeud avec le nombre de reines en prise (représentant la distance du noeud jusqu'au résultat).

## Algorithme A* 
Le graphe de solutions est représenté par toutes les permutations d'un noeud donné. Les noeuds d'un voisin donné sont chacun une permutation aléatoire de ce dit noeud. L'algorithme A* développera chaque voisin étant le plus suceptible de nous rapprocher de la meilleure solution, ce qui permet de développer l'arbre de recherche tout en nous dirigeant à l'aide de la fonction heuristique. 
