package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ProtonDBApp struct {
	BestReportedTier string  `json:"bestReportedTier"`
	Confidence       string  `json:"confidence"`
	Score            float64 `json:"score"`
	Tier             string  `json:"tier"`
	Total            int     `json:"total"`
	TrendingTier     string  `json:"trendingTier"`
}

func SearchProtonDB(steamId string) ProtonDBApp {
	game := ProtonDBApp{}
	res, _ := http.Get(fmt.Sprintf("https://www.protondb.com/api/v1/reports/summaries/%s.json", steamId))
	if res.StatusCode == http.StatusOK {
		dec := json.NewDecoder(res.Body)
		err := dec.Decode(&game)
		if err != nil {
			fmt.Println(err)
		}
	}

	return game
}
