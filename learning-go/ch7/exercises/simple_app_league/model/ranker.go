package model

type Ranker interface {
	Ranking() []string
}
