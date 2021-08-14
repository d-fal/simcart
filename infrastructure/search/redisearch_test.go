package search_test

import (
	"fmt"
	"testing"

	search "simcart/infrastructure/search"

	"github.com/stretchr/testify/assert"
)

func TestRedisearch(t *testing.T) {
	client := search.NewClient("redisearch:6379", "test")

	t.Run("test search a key", func(t *testing.T) {

		client.CreateIndex("test", "actor", 2.0, true)

		doc1 := client.NewDocument("starwars", 1.0, map[string]string{"actor": "someone", "title": "star"})

		doc2 := client.NewDocument("new moon", 1.0, map[string]string{"actor": "someone else", "title": "new moon"})
		client.AddDocument(doc1, doc2)
		docs, total, err := client.Search("cheese", "actor", "title")

		if total > 0 {
			for _, doc := range docs {
				fmt.Println("doc: ", doc.Properties["title"])
			}
		} else {
			assert.Error(t, fmt.Errorf("value not found"))
		}
		assert.NoError(t, err)
	})
}
