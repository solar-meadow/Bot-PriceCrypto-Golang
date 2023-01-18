package service

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/logger"
)

type MessageService struct {
	*http.Client
	botUrl            string
	Offset            int
	AllowedCryptoName (map[string]int)
}

func NewHttpClient(botUrl string) *MessageService {
	return &MessageService{
		&http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				fmt.Println(req.Response.Status)
				fmt.Println("[REDIRECT]")
				return nil
			},
			Transport: &logger.LoggingRoundTripper{
				Logger: os.Stdout,
				Next:   http.DefaultTransport,
			},
			Timeout: time.Second * 30,
		},
		botUrl,
		0,
		nil,
	}
}
