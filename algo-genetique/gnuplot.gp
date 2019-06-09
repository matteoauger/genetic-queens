set terminal png 
set output "../img/result.png"

set xlabel "Taille échiquier (N)"
set ylabel "Temps d'exécution (secondes)"

set title "Comparaison génétique / recuit nreines"

plot "donnees.txt" using 1:2 with lines title "Algo génétique", \
"donnees_recuit.txt" using 1:2 with lines title "Recuit simulé"

