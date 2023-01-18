package main

import (
	"fmt"
	"log"

	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/token"
)

func main() {
	botToken, err := token.BotTokenFromEnv("BOT_TOKEN")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(botToken)
}
