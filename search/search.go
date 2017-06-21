package search

import (
	"errors"
	"log"
	"reflect"

	"encoding/json"
	"golang.org/x/net/context"

	u "github.com/salemark/utils"
	elastic "gopkg.in/olivere/elastic.v5"
)

const (
	elasticHost = "http://localhost:9200"
	indexName   = "salemark"
)

type Search struct {
	Client  *elastic.Client
	Context context.Context
}

func Start() Search {
	context := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(elasticHost))

	u.PanicError(err)

	s := Search{Client: client, Context: context}
	GetInfo(s)

	CreateIndex(s)
	// SingleSeed(s)
	// SeedData(s)

	return s
}

func GetInfo(s Search) {
	info, code, err := s.Client.Ping(elasticHost).Do(s.Context)
	u.PanicError(err)

	log.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)
}

func CreateIndex(s Search) {
	exists, err := s.Client.IndexExists(indexName).Do(s.Context)

	u.PanicError(err)

	if exists {
		log.Printf("Index %s already exists", indexName)
		return
	}
	log.Printf("Index %s not found, creating new...", indexName)

	res, err := s.Client.CreateIndex(indexName).Do(s.Context)

	u.PanicError(err)

	if !res.Acknowledged {
		err := errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
		u.PanicError(err)
	}
}

func SingleSeed(s Search) {
	log.Println("Single Seed")
	brand := Brand{Name: "Apple", Description: "Apple secret"}
	_, err := s.Client.Index().
		Index(indexName).
		Type(brand.Type()).
		Id(brand.Id()).
		BodyJson(brand).
		Do(s.Context)

	u.PanicError(err)
}

func SeedData(s Search) {
	log.Println("Seeding products")
	for _, product := range SeedProducts() {
		_, err := s.Client.Index().
			Index(indexName).
			Type(product.Type()).
			Id(product.Id()).
			BodyJson(product).
			Do(s.Context)

		if err != nil {
			panic(err)
		}
	}

	log.Println("Seeding brands")
	for _, brand := range SeedBrands() {
		_, err := s.Client.Index().
			Index(indexName).
			Type(brand.Type()).
			Id(brand.Id()).
			BodyJson(brand).
			Do(s.Context)

		u.PanicError(err)
	}

	log.Println("Successfully loaded seed")
	_, err := s.Client.Flush().Index(indexName).Do(s.Context)

	u.PanicError(err)
}

func Find(s Search, sType, body string) string {
	matchQuery := elastic.NewMatchPhraseQuery("name", body)

	log.Println("Find Operation Beginns")
	searchResult, err := s.Client.Search().
		Index(indexName). // search in index "twitter"
		Type(sType).
		Query(matchQuery). // specify the query
		// From(0).Size(10). // take documents 0-9
		Pretty(true). // pretty print request and response JSON
		Do(s.Context) // execute
	u.PanicError(err)

	if searchResult.Hits.TotalHits < 1 {
		message := "Nothing found"
		log.Println(message)
		return message
	}

	log.Printf("Found a total of %d brands\n", searchResult.Hits.TotalHits)

	var ttyp Brand
	var jsonArr []string
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if b, ok := item.(Brand); ok {
			bJson, err := json.Marshal(b)
			u.PanicError(err)
			jsonArr = append(jsonArr, string(bJson))
			log.Printf("Brand %s: %s\n", b.Name, b.Description)
		}
	}

	// return searchResult
	jsonResult, err := json.Marshal(jsonArr)
	u.PanicError(err)
	return string(jsonResult)
}
