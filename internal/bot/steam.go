package bot

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

type SteamApp struct {
	AppId string `json:"appid"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Logo  string `json:"logo"`
}

func SearchSteam(input string) []SteamApp {
	games := []SteamApp{}

	res, _ := http.Get(fmt.Sprintf("https://steamcommunity.com/actions/SearchApps/%s", html.EscapeString(input)))
	if res.StatusCode == http.StatusOK {
		dec := json.NewDecoder(res.Body)
		err := dec.Decode(&games)
		if err != nil {
			fmt.Println(err)
		}
	}

	return games
}
