package main

import (
	"log"
	"sync"

	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/logger"
	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/service"
	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/token"
)

const botApi = "https://api.telegram.org/bot"

var (
	AllowedCryptoName   = make(map[string]int)
	AllowedCryptoSymbol = make(map[string]int)
)

func main() {
	botToken, err := token.BotTokenFromEnv("BOT_TOKEN")
	if err != nil {
		log.Println(err)
		return
	}
	botUrl := botApi + botToken
	client := service.NewHttpClient(botUrl)
	client.AllowedCryptoName = AllowedCryptoName
	client.AllowedCryptoSymbol = AllowedCryptoSymbol
	wg := &sync.WaitGroup{}
	// fill maps allowed values for coincap api
	if err = service.FillAllowedValues(&AllowedCryptoName, 1); err != nil {
		log.Println(err)
		return
	}
	if err = service.FillAllowedValues(&AllowedCryptoSymbol, 2); err != nil {
		log.Println(err)
		return
	}
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
