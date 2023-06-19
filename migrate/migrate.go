package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/a4anthony/go-commerce/config"
	"github.com/a4anthony/go-commerce/database"
	"github.com/a4anthony/go-commerce/models"
	"github.com/a4anthony/go-commerce/seeders/seeds"
)

func init() {
	config.LoadENV()
	database.ConnectDb()
}

func main() {
	Migrate()
}

func Migrate() {
	useSeed := flag.Bool("seed", false, "run seeders")
	useRefresh := flag.Bool("refresh", false, "refresh database")
	flag.Parse()

	appModels := []interface{}{
		&models.User{},
	}

	if *useRefresh {
		for _, model := range appModels {
			database.DB.Migrator().DropTable(model)
		}
	}

	for _, model := range appModels {
		fmt.Println(database.DB.Migrator().HasTable(model))
		if !database.DB.Migrator().HasTable(model) {
			database.DB.AutoMigrate(model)
		}
	}

	fmt.Println("Migrations done!")
	if *useSeed {
		fmt.Println("Running seeds...")
		for _, seed := range seeds.All() {
			if err := seed.Run(database.DB); err != nil {
				log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
			}
		}
		fmt.Println("Seeds ran successfully!")
	}

}
