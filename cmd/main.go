package main

import (
	"log"
	"sync"

	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/logger"
	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/service"
	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/token"
)

const botApi = "https://api.telegram.org/bot"

var AllowedCryptoName = make(map[string]int)

func main() {
	botToken, err := token.FromEnv("BOT_TOKEN")
	if err != nil {
		log.Println(err)
		return
	}
	botUrl := botApi + botToken
	client := service.NewHttpClient(botUrl)
	client.AllowedCryptoName = AllowedCryptoName
	wg := &sync.WaitGroup{}
	// fill maps allowed values for coincap api
	service.FillAllowedValues(&AllowedCryptoName)

	botName, err := token.FromEnv("BOT_NAME")
	if err != nil {
		botName = "Bot_Name"
	}
	log.Println("Bot running... name in telegram:", botName)
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
