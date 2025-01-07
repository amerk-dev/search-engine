package main

import (
	"encoding/xml"
	"log"
	"os"
	"searchSystem/internal/dilivery/restapi"
	"searchSystem/internal/index"
	"searchSystem/internal/models"
	"time"
)

func main() {

	start := time.Now()

	docs, err := loadDocuments("enwiki-latest-abstract1.xml")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	// Индексация файлов
	start = time.Now()
	idx := make(index.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	matchedIDs := idx.Search("cat")
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	server := restapi.Server{
		Index: &idx,
	}
	server.StartServer() //запуск сервера

}

func loadDocuments(path string) ([]models.Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := xml.NewDecoder(f)
	dump := struct {
		Documents []models.Document `xml:"doc"`
	}{}
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
