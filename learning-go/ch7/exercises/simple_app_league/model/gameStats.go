package model

type GameStats struct {
	HomeTeamName  string
	AwayTeamName  string
	HomeTeamScore int
	AwayTeamScore int
}

func NewGameStats(homeTeamName string, awayTeamName string, homeTeamScore int, awayTeamScore int) GameStats {
	return GameStats{
		HomeTeamName:  homeTeamName,
		AwayTeamName:  awayTeamName,
		HomeTeamScore: homeTeamScore,
		AwayTeamScore: awayTeamScore,
	}
}

func (gs GameStats) determineWinner() (string, bool) {
	switch {
	case gs.HomeTeamScore > gs.AwayTeamScore:
		return gs.HomeTeamName, true
	case gs.AwayTeamScore > gs.HomeTeamScore:
		return gs.AwayTeamName, true
	default:
		return "", false
	}
}
