package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/meilisearch/meilisearch-go"
)

func main() {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://localhost:7700",
		APIKey: "",
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
