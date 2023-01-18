package logger

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/model"
)

func SaveLog(update *model.Update, wg *sync.WaitGroup) error {
	log.Printf("INFO: Writing file for update ID: %d\n", update.UpdateId)
	filename := fmt.Sprintf("pkg/logger/logs/uid_%d.txt", update.Message.Chat.ChatId)

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(update.Info())
	if err != nil {
		return err
	}
	fmt.Println(update.Info())
	wg.Done()
	return nil
}
