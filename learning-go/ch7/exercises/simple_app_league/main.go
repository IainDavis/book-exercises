package main

import (
	"fmt"
	"os"

	model "iaindavis.dev/learning/learning-go/ch7/exercises/simple_app_league/model"
	"iaindavis.dev/learning/learning-go/ch7/exercises/simple_app_league/output"
)

func main() {
	newSimpleLeague := model.NewSimpleLeague
	newGameStats := model.NewGameStats

	var league model.League = newSimpleLeague(map[string]model.Team{
		"Avalon": {
			Name:    "Avalon",
			Players: []string{"Vikram", "Savitha", "Ratheesh"},
		},
		"Bronte": {
			Name:    "Bronte",
			Players: []string{"Vikrant", "Vicki"},
		},
		"Cronulla": {
			Name:    "Cronulla",
			Players: []string{"Chris", "Iain", "Swati"},
		},
		"Delray": {
			Name:    "Delray",
			Players: []string{"Vijay", "Anjana"},
		},
	})

	league.MatchResult(newGameStats("Avalon", "Cronulla", 3, 8))
	league.MatchResult(newGameStats("Bronte", "Cronulla", 2, 4))
	league.MatchResult(newGameStats("Avalon", "Bronte", 5, 2))
	league.MatchResult(newGameStats("Athletics", "Giants", 8, 1))

	fmt.Println("Wins:", model.SimpleLeague(league.(model.SimpleLeague)).Wins)

	output.RankPrinter(league, os.Stdout)

}
