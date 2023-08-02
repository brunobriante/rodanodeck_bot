package bot

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strconv"
)

type SteamApp struct {
	AppId string `json:"appid"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Logo  string `json:"logo"`
}

func SearchSteam(input string) []SteamApp {
	games := []SteamApp{}

	res, err := http.Get(fmt.Sprintf("https://steamcommunity.com/actions/SearchApps/%s", html.EscapeString(input)))
	if err != nil {
		fmt.Print("[SearchSteam-Request]")
		fmt.Println(err)
	}
	if res.StatusCode == http.StatusOK {
		dec := json.NewDecoder(res.Body)
		err := dec.Decode(&games)
		if err != nil {
			fmt.Print("[SearchSteam-Decode]")
			fmt.Println(err)
		}
	} else {
		fmt.Println(("[SearchSteam-Request] StatusCode Failed: " + strconv.Itoa(res.StatusCode)))
	}

	return games
}
