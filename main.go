package main

import (
	"fmt"
	"os"

	"github.com/archway-network/valuter/api"
	"github.com/archway-network/valuter/database"
)

/*--------------*/

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	database.DB = database.New(database.Postgres, psqlconn)
	defer database.DB.Close()

	api.ListenAndServeHTTP(os.Getenv("SERVING_ADDR"))

}
