package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
	"github.com/meilisearch/meilisearch-go"
)

func main() {
	loadEnv()

	index := os.Args[1]

	jsonFile, errorJson := os.Open(fmt.Sprintf("./data/%s.json", index))
	defer jsonFile.Close()
	if errorJson != nil {
		panic(errorJson)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var data []map[string]interface{}
	json.Unmarshal(byteValue, &data)

	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   os.Getenv("MEILI_HOST"),
		APIKey: os.Getenv("MEILI_MASTER_KEY"),
	})

	_, err := client.Index(index).AddDocuments(data)
	if err != nil {
		panic(err)
	}
}

func loadEnv() {
	error := godotenv.Load(".env")

	if error != nil {
		panic(error)
	}
}
