package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var config *Config

func init() {
	var err error
	config, err = getEnvConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("error setting up config %s", err.Error()))
	}
}

func main() {
	db, err := sqlx.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%v)/",
			config.DBUser,
			config.DBPass,
			config.DBHost,
			config.DBPort,
		),
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("error connecting to db %s", err.Error()))
	}

	fmt.Println("setting up store")
	store := NewStore(db)

	fmt.Println("setting up server")
	server := NewServer(store)

	fmt.Println("setting up router")
	router := configureRouter(server)

	fmt.Printf("listening on port %s", config.HTTPPort)
	log.Fatal(http.ListenAndServe(":"+config.HTTPPort, router))
}
