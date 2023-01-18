package logger

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/model"
)

func SaveLog(update *model.Update, wg *sync.WaitGroup) error {
	path := "pkg/logger/logs/"
	filename := fmt.Sprintf(path+"uid_%d.txt", update.Message.Chat.ChatId)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	file = append(file, []byte(update.Info())...)

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}
	info := (update.Info())
	fmt.Println()

	fmt.Print(info)
	for i := range info {
		fmt.Print("-")
		if i == len(info)-1 {
			fmt.Print("-")
			fmt.Println()
		}
	}
	wg.Done()
	return nil
}
