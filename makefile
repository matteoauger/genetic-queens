genetic_queens: 
	go build -o bin/genetic_queens.x algo-genetique/src/main.go

recuit_simule:
	go build -o bin/recuit_simule.x recuit-simule/main.go

clean : 
	rm bin/*