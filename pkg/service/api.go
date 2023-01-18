package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type AssetsResponse struct {
	Data      []AssetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type AssetData struct {
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

func (m *MessageService) GetInfoByID(id string) (AssetData, error) {
	resp, err := m.Get("https://api.coincap.io/v2/assets/" + id)
	if err != nil {
		return AssetData{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AssetData{}, err
	}
	var response AssetData
	if err = json.Unmarshal(body, &response); err != nil {
		return AssetData{}, err
	}

	return response, nil
}

func (m *MessageService) GetAllInfo() (AssetsResponse, error) {
	resp, err := m.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		return AssetsResponse{}, err
	}
	defer resp.Body.Close()
	fmt.Println("response status:", resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AssetsResponse{}, err
	}
	var response AssetsResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return AssetsResponse{}, err
	}
	return response, nil
}
