set terminal png 
set output "result.png"

set xlabel "itération"
set ylabel "qualité et température"

set title "recuit simulé nreines"

plot "results.txt" using 1:2 with lines title "température", \
"results.txt" using 1:3 with lines title "qualité"
