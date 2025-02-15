package main

import (
	"coincap/coincap"
	"fmt"
	"log"
)

func main() {
	coincapClient, err := coincap.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	assets, err := coincapClient.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	for _, asset := range assets {
		fmt.Println(asset.Info())
	}

	asset, err := coincapClient.GetAsset("bitcoin")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(asset.Info())
}
