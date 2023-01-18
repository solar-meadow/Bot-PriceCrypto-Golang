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
	result := "Write the ID(ex: bitcoin) of the cryptocurrency to get information\n\n\n Write \"list\" to view all cryptocurrency ID"
	if update.Message.Text == "list" {
		data, err := m.GetAllInfo()
		if err != nil {
			return err
		}
		result = ""
		for _, v := range data.Data {
			result += fmt.Sprintf("Rank: [%s] - ID: [%s]", v.Rank, v.ID)
			result += "\n"
		}
		botMessage.Text = result
	} else if _, ok := m.AllowedCryptoName[strings.ToLower(update.Message.Text)]; ok {
		data, err := m.GetInfoByID(strings.ToLower(update.Message.Text))
		if err != nil {
			return err
		}
		result = fmt.Sprintf("Name: <%s> \nRank: <%s> \nPriceUSD$: <%s> ", data.Data.Name, data.Data.Rank, data.Data.PriceUSD)
	}
	botMessage.Text = result

	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	if _, err = m.Post(m.botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf)); err != nil {
		return err
	}
	return nil
}
