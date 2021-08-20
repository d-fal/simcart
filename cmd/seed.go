package cmd

import (
	"fmt"
	"html"
	"log"
	"reflect"
	"simcart/api/pb/commonpb"
	"simcart/app/scaffold"
	"simcart/config"
	product_entity "simcart/domain/product/entity"
	"simcart/infrastructure/postgres"
	"simcart/infrastructure/search"
	"simcart/pkg/model"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var (
	seedCmd = &cobra.Command{
		Use:   "seed",
		Short: "generates seeds for dbs",
		Run:   Runner.seed,
	}
)

func init() {

}

func (c *command) seed(cmd *cobra.Command, args []string) {
	// test
	setup := s.Start(systemwideContext, false,
		scaffold.WithPostgres(),
		scaffold.WithRedisearch(),
	)
	if err := setup(); err != nil {
		panic(fmt.Errorf("error commissioning the clients. %v\n", aurora.Red(err)))
	}

	tx, err := postgres.Storage.Transaction()

	if err := tx.Begin(); err != nil {
		log.Fatal("cannot begin database transaction")
	}
	defer tx.Close()

	if err != nil {
		panic(fmt.Errorf("error instantiating db%v\n", aurora.Red(err)))
	}

	insert := func(m interface{}) {

		if _, err := tx.Get().Model(m).OnConflict("(id) do update").Insert(); err != nil {
			fmt.Printf("%v cannot insert %v %v\n", aurora.White(html.UnescapeString("&#x274C;")), aurora.Yellow(reflect.TypeOf(m).Elem()), err)
		}
	}

	products := []*product_entity.Product{
		{
			Model:    model.Model{Id: 1},
			Title:    "T-Shirt",
			Sku:      "100200200",
			Category: commonpb.Category_Apparel,
			Price:    100,
			Currency: commonpb.Currency_EUR,
			Descriptions: map[string]string{
				"size":     "XL",
				"color":    "red",
				"material": "silk",
				"origin":   "Turkey",
			},
		},
		{
			Model:    model.Model{Id: 2},
			Title:    "LG TV",
			Category: commonpb.Category_Appliance,
			Sku:      "100200300",
			Price:    80,
			Currency: commonpb.Currency_EUR,
			Descriptions: map[string]string{
				"size":   "53 inches",
				"panel":  "IPS",
				"origin": "S.Korea",
				"color":  "gray",
			},
		},
		{
			Model:    model.Model{Id: 3},
			Title:    "Cheese 250 gr",
			Sku:      "100200400",
			Price:    2,
			Currency: commonpb.Currency_GBP,
			Category: commonpb.Category_FMCG,
			Descriptions: map[string]string{
				"weight": "250gr",
				"type":   "bulgarian",
				"shape":  "rounded",
			},
		},

		{
			Model:    model.Model{Id: 4},
			Title:    "Cheese 300 gr",
			Category: commonpb.Category_FMCG,
			Sku:      "100200500",
			Price:    3,
			Currency: commonpb.Currency_GBP,
			Descriptions: map[string]string{
				"weight":  "300gr",
				"type":    "french",
				"shape":   "square",
				"flavour": "salty",
			},
		},

		{
			Model:    model.Model{Id: 5},
			Title:    "Refregerator",
			Sku:      "100200600",
			Price:    2000,
			Currency: commonpb.Currency_EUR,
			Category: commonpb.Category_Appliance,
			Descriptions: map[string]string{
				"weight": "40kg",
				"origin": "Taiwan",
				"size":   "7ft",
				"color":  "silver",
			},
		},
		{
			Model:    model.Model{Id: 6},
			Title:    "Samsung",
			Sku:      "100200700",
			Price:    1100,
			Currency: commonpb.Currency_EUR,
			Category: commonpb.Category_Appliance,
			Descriptions: map[string]string{
				"weight": "700gr",
				"origin": "S.Korea",
				"size":   "2.5 inches",
				"color":  "black",
			},
		},
		{
			Model:    model.Model{Id: 7},
			Title:    "Apple iPhone 13",
			Sku:      "100200800",
			Price:    900,
			Currency: commonpb.Currency_EUR,
			Category: commonpb.Category_Appliance,
			Descriptions: map[string]string{
				"weight": "700gr",
				"origin": "USA",
				"size":   "2.5 inches",
				"color":  "golden",
			},
		},
	}

	insert(&products)

	// create account

	if err := tx.Commit(); err != nil {
		log.Fatal("cannot commit changes to db")
	}
	// insert in reids

	client := search.NewClient(config.GetAppConfig().ClientRedisearch().Addr(), "indexer")

	if err := client.Flush(); err != nil {
		log.Printf("Cannot flush all the indices in search db: why? %v\n", err)
	}

	schema := redisearch.NewSchema(
		redisearch.DefaultOptions).AddField(redisearch.NewTextField("color")).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 3, Sortable: true})).
		AddField(redisearch.NewNumericField("price")).
		AddField(redisearch.NewTextField("weight")).
		AddField(redisearch.NewTextField("size")).
		AddField(redisearch.NewTextField("cat")).
		AddField(redisearch.NewTextField("brand"))

	if err := client.CreateIndexFromSchema(schema); err != nil {
		log.Fatalf("cannot create index: %v\n", err)
	}

	for _, p := range products {
		serialized, _ := p.Marshal()
		doc := client.NewDocument(p.Sku, 1, p.Descriptions)
		doc.Set("date", p.CreatedAt.Unix()).
			Set("title", p.Title).
			Set("cat", p.Category.String()).
			Set("price", p.Price).
			Set("currency", p.Currency).
			Set("sku", p.Sku)

		// storing the whole object for seamless retrieval
		doc.SetPayload(serialized)

		if err := client.AddDocument(doc); err != nil {
			log.Fatalf("cannot create document")
		}
	}

	fmt.Printf("\n%v Seeding finished.\n", aurora.Green(html.UnescapeString("&#x2705;")))
}
