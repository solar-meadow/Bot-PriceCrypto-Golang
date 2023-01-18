package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/with-insomnia/Bot-PriceCrypto-Golang/pkg/model"
)

func (m *MessageService) GetUpdates() ([]model.Update, error) {
	resp, err := http.Get(m.botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(m.Offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse model.RestResponse
	if err = json.Unmarshal(body, &restResponse); err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}
