package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/joho/godotenv"
	"github.com/meilisearch/meilisearch-go"
)

func main() {
	loadEnv()

	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   os.Getenv("MEILI_HOST"),
		APIKey: os.Getenv("MEILI_MASTER_KEY"),
	})

	jsonFile, _ := os.Open("./data/movies.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var data []map[string]interface{}
	json.Unmarshal(byteValue, &data)

	_, err := client.Index("movies").AddDocuments(data)
	if err != nil {
		panic(err)
	}
}

func loadEnv() {
	errEnv := godotenv.Load(".env")

	if errEnv != nil {
		panic("Fail loading .env file")
	}
}
