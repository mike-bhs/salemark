package search

import (
	// "encoding/json"
	"errors"
	"log"
	"reflect"

	"golang.org/x/net/context"

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

	if err != nil {
		panic(err)
	}

	s := Search{Client: client, Context: context}
	GetInfo(s)

	CreateIndex(s)
	// SingleSeed(s)
	// SeedData(s)

	return s
}

func GetInfo(s Search) {
	info, code, err := s.Client.Ping(elasticHost).Do(s.Context)

	if err != nil {
		panic(err)
	}

	log.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)
}

func CreateIndex(s Search) {
	exists, err := s.Client.IndexExists(indexName).Do(s.Context)

	if err != nil {
		panic(err)
	}

	if exists {
		log.Printf("Index %s already exists", indexName)
		return
	}
	log.Printf("Index %s not found, creating new...", indexName)

	res, err := s.Client.CreateIndex(indexName).Do(s.Context)

	if err != nil {
		panic(err)
	}

	if !res.Acknowledged {
		err := errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
		panic(err)
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

	if err != nil {
		panic(err)
	}
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

		if err != nil {
			panic(err)
		}
	}

	log.Println("Successfully loaded seed")
	_, err := s.Client.Flush().Index(indexName).Do(s.Context)

	if err != nil {
		panic(err)
	}
}

func Find(s Search) {
	termQuery := elastic.NewMatchPhraseQuery("name", "apple")
	log.Println("Find Operation Beginns")
	searchResult, err := s.Client.Search().
		Index(indexName). // search in index "twitter"
		Type("brands").
		Query(termQuery). // specify the query
		From(0).Size(10). // take documents 0-9
		Pretty(true).     // pretty print request and response JSON
		Do(s.Context)     // execute
	if err != nil {
		panic(err)
	}

	if searchResult.Hits.TotalHits < 1 {
		log.Println("Nothing found")
		return
	}
	log.Printf("Found a total of %d brands\n", searchResult.Hits.TotalHits)

	var ttyp Brand
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if b, ok := item.(Brand); ok {
			log.Printf("Brand %s: %s\n", b.Name, b.Description)
		}
	}
}
