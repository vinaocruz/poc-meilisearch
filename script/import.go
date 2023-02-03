package main

import (
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
	"github.com/meilisearch/meilisearch-go"
)

func main() {
	loadEnv()

	index := os.Args[1]

	jsonFile, error := os.Open(fmt.Sprintf("./data/%s.ndjson", index))
	defer jsonFile.Close()
	if error != nil {
		panic(error)
	}

	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   os.Getenv("MEILI_HOST"),
		APIKey: os.Getenv("MEILI_MASTER_KEY"),
	})

	_, err := client.Index(index).AddDocumentsNdjsonFromReader(io.Reader(jsonFile))
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
