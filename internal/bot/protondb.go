package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	res, err := http.Get(fmt.Sprintf("https://www.protondb.com/api/v1/reports/summaries/%s.json", steamId))
	if err != nil {
		fmt.Print("[SearchProtonDB-Request]")
		fmt.Println(err)
	}

	if res.StatusCode == http.StatusOK {
		dec := json.NewDecoder(res.Body)
		err := dec.Decode(&game)
		if err != nil {
			fmt.Print("[SearchProtonDB-Decode]")
			fmt.Println(err)
		}
	} else {
		fmt.Println(("[SearchProtonDB-Request] StatusCode Failed: " + strconv.Itoa(res.StatusCode)))
	}

	return game
}
