package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/model"
)

func (m *MessageService) Respond(update model.Update) error {
	var botMessage model.BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	result := ""
	if update.Message.Text == "all" {
		data, err := m.GetAllInfo()
		if err != nil {
			return err
		}
		result = ""
		for _, v := range data.Data {
			result += fmt.Sprintf("Name: [%s] | Rank: [%s] | PriceUSD: [%s]", v.Name, v.Rank, v.PriceUSD)
			result += "\n"
		}
		botMessage.Text = result
	} else if _, ok := m.AllowedCryptoName[strings.ToUpper(update.Message.Text)]; ok {
		data, err := m.GetInfoByID(strings.ToLower(update.Message.Text))
		if err != nil {
			return err
		}
		result = fmt.Sprintf("Name: [%s] | Rank: [%s] | PriceUSD: [%s]", data.Name, data.Rank, data.PriceUSD)
	} else if _, ok := m.AllowedCryptoSymbol[strings.ToUpper(update.Message.Text)]; ok {
		index := m.AllowedCryptoSymbol[strings.ToUpper(update.Message.Text)]
		for value := range m.AllowedCryptoName {
			if m.AllowedCryptoName[value] == index {
				data, err := m.GetInfoByID(strings.Title(update.Message.Text))
				if err != nil {
					return err
				}
				result = fmt.Sprintf("Name: [%s] | Rank: [%s] | PriceUSD: [%s]", data.Name, data.Rank, data.PriceUSD)
				break
			}
		}
	}
	botMessage.Text = result
	if len(botMessage.Text) == 0 {
		botMessage.Text = "Write the name(ex: Bitcoin) or symbol(ex: BTC) of the cryptocurrency to get information, or write \"all\" to get info about top cryptocurrencies."
	}
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	if _, err = m.Post(m.botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf)); err != nil {
		return err
	}
	return nil
}
