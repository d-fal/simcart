package search

import (
	"fmt"
	"html"
	"sync"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/logrusorgru/aurora"
)

var (
	once     sync.Once
	instance *impl
)

type (
	SearchPrototype interface {
		GetRedisearchClient() *redisearch.Client
		CreateIndex(schema, index string, weight float32, sortable bool) error
		CreateIndexFromSchema(schema *redisearch.Schema) error
		NewDocument(id string, score float32, fields map[string]string) redisearch.Document
		AddDocument(doc ...redisearch.Document) error
		Search(q string, fields ...string) ([]redisearch.Document, int, error)
		Flush() error
	}

	impl struct {
		client *redisearch.Client
	}
)

func NewClient(addr, name string) SearchPrototype {

	once.Do(func() {
		instance = new(impl)
		instance.client = redisearch.NewClient(addr, name)
	})
	return instance
}

func GetClient() (SearchPrototype, error) {
	if instance == nil {
		return nil, fmt.Errorf("redisearch instance not available")
	}
	return instance, nil
}

// creates a single indexed schema
// for more complicated indicies, invoke GetRedisearchClient
func (i *impl) CreateIndex(schema, index string, weight float32, sortable bool) error {
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField(schema)).
		AddField(redisearch.NewTextFieldOptions(index, redisearch.TextFieldOptions{
			Weight: weight, Sortable: true,
		}))

	if err := i.client.CreateIndex(sc); err != nil {
		return err
	}
	return nil
}
func (i *impl) Search(q string, fields ...string) ([]redisearch.Document, int, error) {

	return i.client.Search(redisearch.NewQuery(q).
		Limit(0, 10).SetReturnFields(fields...).SetFlags(redisearch.QueryWithPayloads))

}
func (i *impl) AddDocument(doc ...redisearch.Document) error {
	return i.client.Index(doc...)
}

func (i *impl) NewDocument(id string, score float32, fields map[string]string) redisearch.Document {

	doc := redisearch.NewDocument(id, score)

	for key, f := range fields {
		fmt.Printf("%v sku: %v \t key: %v , %v\n",
			aurora.Green(html.UnescapeString("&#x2705;")), aurora.Red(id), key, f)
		doc.Set(key, f)
	}
	doc = doc.Set("date", time.Now().Unix())

	return doc

}

func (i *impl) GetRedisearchClient() *redisearch.Client {
	return i.client
}

func (i *impl) Flush() error {
	return i.client.Drop()
}

func (i *impl) CreateIndexFromSchema(schema *redisearch.Schema) error {
	return i.client.CreateIndex(schema)
}
