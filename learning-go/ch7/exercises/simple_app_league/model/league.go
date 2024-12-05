package model

import (
	"fmt"
	"sort"
)

type League interface {
	MatchResult(gs GameStats)
	Ranking() []string
}

type SimpleLeague struct {
	Teams map[string]Team
	Wins  map[string]int
}

func NewSimpleLeague(teams map[string]Team) SimpleLeague {
	return SimpleLeague{
		Teams: teams,
		Wins:  make(map[string]int, len(teams)),
	}
}

func (sl SimpleLeague) MatchResult(gs GameStats) {
	_, th_ok := sl.Teams[gs.HomeTeamName]
	_, ta_ok := sl.Teams[gs.AwayTeamName]
	msg := ""
	switch {
	case !th_ok:
		msg += fmt.Sprintf("Unrecognized team: %v\n", gs.HomeTeamName)
		fallthrough
	case !ta_ok:
		msg += fmt.Sprintf("Unrecognized team: %v\n", gs.AwayTeamName)
		fallthrough
	case !ta_ok || !th_ok:
		fmt.Println(msg)
		return
	}

	winner, ok := gs.determineWinner()
	if !ok {
		fmt.Printf("No winner in the match between %s and %s/n", gs.HomeTeamName, gs.AwayTeamName)
		return
	}
	sl.Wins[winner]++
}

func (sl SimpleLeague) Ranking() []string {
	rankedNames := make([]string, 0, len(sl.Teams))
	for t := range sl.Teams {
		rankedNames = append(rankedNames, t)
	}
	sort.Slice(rankedNames, func(nameA, nameB int) bool {
		return sl.Wins[rankedNames[nameA]] > sl.Wins[rankedNames[nameB]]
	})
	return rankedNames
}
