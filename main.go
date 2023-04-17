package main

import belajar_test "github.com/imron16/go-hello/belajar"

func main() {
	// Create a world cup
	worldCup := belajar_test.WorldCup{Orgnizer: "FIFA", Country: "Qatar", CountryPlayer: 4}
	worldCup.StartGame()

	// Create a tournament world cup
	tournamentWorldCup := belajar_test.TournamentWorldCup{WorldCup: belajar_test.WorldCup{Orgnizer: "FIFA", Country: "QATAR", CountryPlayer: 4}, Winner: "Argentina", Prizes: "$42,000,000"}
	tournamentWorldCup.StartTournament()
}
