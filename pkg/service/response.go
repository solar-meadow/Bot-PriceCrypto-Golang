package service

import (
	"bytes"
	"encoding/json"

	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/model"
)

func (m *MessageService) Respond(update model.Update) error {
	var botMessage model.BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	switch update.Message.Text {
	case "/start":
		botMessage.Text = "Hello this is test"
	default:
		botMessage.Text = "Okay fine"
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
