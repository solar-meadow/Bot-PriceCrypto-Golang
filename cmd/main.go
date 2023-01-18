package main

import (
	"log"
	"sync"

	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/logger"
	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/service"
	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/token"
)

const botApi = "https://api.telegram.org/bot"

func main() {
	botToken, err := token.BotTokenFromEnv("BOT_TOKEN")
	if err != nil {
		log.Println(err)
		return
	}
	botUrl := botApi + botToken
	client := service.NewHttpClient(botUrl)
	wg := &sync.WaitGroup{}
	for {
		updates, err := client.GetUpdates()
		if err != nil {
			log.Println(err)
			return
		}
		for _, update := range updates {
			err = client.Respond(update)
			if err != nil {
				log.Println(err)
				return
			}
			client.Offset = update.UpdateId + 1
			wg.Add(1)
			go logger.SaveLog(&update, wg)
			wg.Wait()
		}
	}
}
