package bot

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func CreateResult(app SteamApp) tele.Result {
	sdhq := SearchSDHQ(app.AppId)
	protondb := SearchProtonDB(app.AppId)
	gameData := GameData{app, protondb, sdhq}

	var msg bytes.Buffer
	// tlp, err := template.ParseFiles("message.tmpl")
	tlp, err := template.New("message.tmpl").Funcs(template.FuncMap{"ToUpper": strings.ToUpper, "ToPercentage": RatingPercentage, "ToStars": RatingStar, "BoolEmoji": BoolEmoji}).ParseFiles("message.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	tlp.Execute(&msg, gameData)

	fmt.Println(RatingStar(sdhq.Acf.SdhqRating))

	result := &tele.PhotoResult{
		URL:      app.Logo,
		ThumbURL: app.Logo,
		Width:    184,
		Height:   69,
	}

	result.SetContent(&tele.InputTextMessageContent{
		Text:           msg.String(),
		DisablePreview: true,
		ParseMode:      tele.ModeHTML,
	})

	return result
}

func RatingPercentage(f float64) string {
	i := int(f * 100)
	return strconv.Itoa(i)
}

func RatingStar(i int) string {
	stars := ""

	for j := 1; j < i; j++ {
		stars = stars + "⭐"
	}

	return stars
}

func BoolEmoji(b bool) string {
	if b {
		return "✅"
	} else {
		return "❎"
	}
}
