package cmd

import (
	"fmt"
	"html"

	"reflect"
	"simcart/clients/postgres"
	cart_entity "simcart/domain/cart/entity"
	"simcart/domain/product/entity"

	"github.com/go-pg/pg/v10/orm"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var (
	migrateCMD = &cobra.Command{
		Use:     "migrate",
		Long:    "migrate database structures. This will migrate tables",
		Aliases: []string{"m"},
		// Run:              Runner.migrate,
		PersistentPreRun: Runner.migrate,
		TraverseChildren: true,
	}

	migrateCreateCMD = &cobra.Command{
		Use:     "create",
		Long:    "Create the given extensions. in the argument ",
		Example: "Example: create --extensions [ pgcrypto | uuid-ossp ]",
		Run:     Runner.extensions,
	}

	migrateCreateModelsCMD = &cobra.Command{
		Use:              "model",
		Long:             "Create models in the connected databases",
		Run:              Runner.createModels,
		TraverseChildren: true,
	}

	extension []string
)

func init() {

	migrateCMD.AddCommand(migrateCreateCMD)
	migrateCreateCMD.AddCommand(migrateCreateModelsCMD)
	migrateCreateCMD.PersistentFlags().StringSliceVar(
		&extension,
		"extensions",
		nil,
		"create extensions [slice of string], example --extensions pgcrypto,uuid-ossp")

}

// migrate database with fake data
func (c *command) migrate(cmd *cobra.Command, args []string) {

	// 1- create database if not exists
	// we need configs be loaded at this place
	// however it isn't
	if err := s.Hibernate(systemwideContext); err != nil {
		panic(fmt.Errorf("error commissioning the clients. %v\n", aurora.Red(err)))
	}

	_, err := postgres.Storage.Get()
	if err != nil {
		panic(fmt.Errorf("cannot connect to db %v\n", aurora.Red(err)))
	}

}

// migrate database with fake data
func (c *command) extensions(cmd *cobra.Command, args []string) {

	db, err := postgres.Storage.Get()

	if err != nil {
		panic(fmt.Errorf("error creating db connection\n%v\n\n", aurora.Red(err)))
	}

	for _, ext := range extension {
		if _, err := db.Exec(fmt.Sprintf("create extension if not exists \"%s\"", ext)); err != nil {
			panic(fmt.Errorf("cannot create extension: %v\n", aurora.Red(err)))
		}
		fmt.Printf("%v\textension %v\n", aurora.Green("[Created]"), ext)
	}

}

func (c *command) createModels(cmd *cobra.Command, args []string) {

	c.extensions(cmd, args)

	db, err := postgres.Storage.Get()

	if err != nil {
		panic(fmt.Errorf("error creating db connection\n%v\n\n", aurora.Red(err)))
	}

	// create composites
	// // create models
	composites := map[string]interface{}{}
	for key, m := range composites {
		err := db.Model(m).CreateComposite(&orm.CreateCompositeOptions{})
		if err != nil {
			fmt.Printf("%v Create composite %v why? %v\n", aurora.Red(html.UnescapeString("&#x274C;")), aurora.Yellow(key), err)
		}
		fmt.Printf("%v composite %v is ready.\n", aurora.Green(html.UnescapeString("&#x2705;")), aurora.Yellow(key))
	}

	// // create models
	models := []interface{}{
		(*entity.Product)(nil),
		(*cart_entity.Cart)(nil),
	}

	for _, m := range models {

		err := db.Model(m).CreateTable(&orm.CreateTableOptions{
			Temp:          false,
			FKConstraints: true,
			IfNotExists:   true,
		})

		if err != nil {
			fmt.Printf("%v Create %v %v\n",
				aurora.White(html.UnescapeString("&#x274C;")),
				aurora.Yellow(reflect.TypeOf(m).Elem()), err)
		}

		fmt.Printf("%s table %v is ready.\n", aurora.Green(html.UnescapeString("&#x2705;")), aurora.Yellow(reflect.TypeOf(m).Elem()))
	}

	fmt.Printf("\n%v Models created.\n", aurora.Green(html.UnescapeString("&#x2705;")))
}
