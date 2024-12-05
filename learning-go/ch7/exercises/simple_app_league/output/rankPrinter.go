package output

import (
	"fmt"
	"io"

	"iaindavis.dev/learning/learning-go/ch7/exercises/simple_app_league/model"
)

func RankPrinter(ranker model.Ranker, writer io.Writer) {
	ranks := ranker.Ranking()
	for i, t := range ranks {
		fmt.Println(i+1, "--", t)
	}
}
