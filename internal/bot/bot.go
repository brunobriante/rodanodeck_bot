package bot

import (
	"fmt"
	"strconv"
	"time"

	"github.com/brunobriante/rodanodeck_bot/configs"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

type GameData struct {
	Steam    SteamApp
	ProtonDB ProtonDBApp
	SDHQ     SDHQApp
}

func New() (*tele.Bot, error) {
	configs.InitConfigs()

	pref := tele.Settings{
		Token:  configs.Configs.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func Start(b *tele.Bot) {
	fmt.Println("Starting Bot")
	b.Use(middleware.Logger())
	b.Use(middleware.AutoRespond())

	b.Handle(tele.OnQuery, func(c tele.Context) error {
		games := SearchSteam(c.Query().Text)
		results := QueryResults(games)

		return c.Answer(&tele.QueryResponse{
			Results: results,
		})
	})

	b.Start()
}

func QueryResults(games []SteamApp) tele.Results {
	results := make(tele.Results, len(games))
	for i, app := range games {

		result := CreateResult(app)

		results[i] = result
		results[i].SetResultID(strconv.Itoa(i))
	}

	return results
}
