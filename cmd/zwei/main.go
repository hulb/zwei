package main

import (
	"log"

	"github.com/hanguofeng/gocaptcha"
	"github.com/jqs7/zwei/biz"
	"github.com/jqs7/zwei/bot/tg"
	"github.com/jqs7/zwei/db"
	"github.com/jqs7/zwei/model"
	"github.com/jqs7/zwei/scheduler"
)

func main() {
	filterConfig := new(gocaptcha.FilterConfig)
	filterConfig.Init()
	filterConfig.Filters = []string{
		gocaptcha.IMAGE_FILTER_NOISE_LINE,
		gocaptcha.IMAGE_FILTER_NOISE_POINT,
		gocaptcha.IMAGE_FILTER_STRIKE,
	}
	for _, v := range filterConfig.Filters {
		filterConfigGroup := new(gocaptcha.FilterConfigGroup)
		filterConfigGroup.Init()
		filterConfigGroup.SetItem("Num", "180")
		filterConfig.SetGroup(v, filterConfigGroup)
	}
	idiomCount, err := db.Instance().Model(new(model.Idiom)).Count()
	if err != nil {
		log.Fatalln(err)
	}
	bot := tg.NewBot(
		"",
		biz.Handler{
			ImageConfig: &gocaptcha.ImageConfig{
				Width:    320,
				Height:   100,
				FontSize: 80,
				FontFiles: []string{
					"fonts/STFANGSO.ttf",
					"fonts/STHEITI.ttf",
					"fonts/STXIHEI.ttf",
				},
			},
			IdiomCount:         idiomCount,
			ImageFilterManager: gocaptcha.CreateImageFilterManagerByConfig(filterConfig),
		},
		tg.Debug(),
	)
	go func() {
		err := scheduler.New(db.Instance(), bot.BotAPI).Run()
		log.Panic(err)
	}()
	bot.Run()
}
