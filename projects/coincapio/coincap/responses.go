package coincap

import "fmt"

type assetsResponse struct {
	Assets    []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

type assetResponse struct {
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type Asset struct {
	ID           string `json:"id"`
	Rank         string `json:"rank"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Supply       string `json:"supply"`
	MaxSupply    string `json:"maxSupply"`
	MarketCapUsd string `json:"marketCapUsd"`
}

func (d Asset) Info() string {
	return fmt.Sprintf("ID: %s, Rank: %s, Symbol: %s, Name: %s, Supply: %s, MaxSupply: %s, MarketCapUsd: %s", d.ID, d.Rank, d.Symbol, d.Name, d.Supply, d.MaxSupply, d.MarketCapUsd)
}
