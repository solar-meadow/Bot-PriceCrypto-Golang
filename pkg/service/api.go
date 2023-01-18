package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type assetsResponse struct {
	Data      []assetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type assetResponse struct {
	Data      assetData `json:"data"`
	Timestamp int64     `json:"timestamp"`
}

type assetData struct {
	ID        string `json:"id"`
	Rank      string `json:"rank"`
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	Supply    string `json:"supply"`
	MaxSupply string `json:"maxSupply"`
	Market    string `json:"marketCapUsd"`
	Volume    string `json:"volumeUsd24Hr"`
	PriceUSD  string `json:"priceUsd"`
	Change    string `json:"changePercent24Hr"`
	Vwap24Hr  string `json:"vwap24Hr"`
	Explorer  string `json:"explorer"`
}

func (m *MessageService) GetInfoByID(id string) (assetResponse, error) {
	resp, err := m.Get("https://api.coincap.io/v2/assets/" + id)
	if err != nil {
		return assetResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return assetResponse{}, err
	}
	var response assetResponse

	if err = json.Unmarshal(body, &response); err != nil {
		return assetResponse{}, err
	}

	return response, nil
}

func (m *MessageService) GetAllInfo() (assetsResponse, error) {
	resp, err := m.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		return assetsResponse{}, err
	}
	defer resp.Body.Close()
	fmt.Println("response status:", resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return assetsResponse{}, err
	}
	var response assetsResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return assetsResponse{}, err
	}
	for _, v := range response.Data {
		fmt.Printf("\"%s\",\n", v.ID)
	}
	return response, nil
}
