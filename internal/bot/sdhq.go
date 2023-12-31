package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type SDHQApp struct {
	Link string `json:"link"`
	Acf  struct {
		BestOnDeck           bool `json:"best_on_deck"`
		SdhqRating           int  `json:"sdhq_rating"`
		SdhqRatingCategories struct {
			Performance    int    `json:"performance"`
			Visuals        int    `json:"visuals"`
			Stability      int    `json:"stability"`
			Controls       int    `json:"controls"`
			Battery        int    `json:"battery"`
			ScoreBreakdown string `json:"score_breakdown"`
		} `json:"sdhq_rating_categories"`
	} `json:"acf"`
}

func SearchSDHQ(steamId string) SDHQApp {
	game := []SDHQApp{}
	res, err := http.Get(fmt.Sprintf("https://steamdeckhq.com/wp-json/wp/v2/game-reviews/?meta_key=steam_app_id&meta_value=%s", steamId))
	if err != nil {
		fmt.Print("[SearchSDHQ-Request]")
		fmt.Println(err)
	}

	if res.StatusCode == http.StatusOK {
		dec := json.NewDecoder(res.Body)
		err := dec.Decode(&game)
		if err != nil {
			fmt.Println(err)
			fmt.Print("[SearchSDHQ-Decode]")
			fmt.Println(err)
		}
	} else {
		fmt.Println(("[SearchSDHQ-Request] StatusCode Failed: " + strconv.Itoa(res.StatusCode)))
	}

	if len(game) > 0 {
		fmt.Println("0 games found on SDHQ for " + steamId)
		return game[len(game)-1]
	}

	return SDHQApp{}
}
