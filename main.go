package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/archway-network/cosmologger/database"
	"github.com/archway-network/valuter/api"
	"github.com/archway-network/valuter/cmd"
	"github.com/archway-network/valuter/simplecache"
)

/*--------------*/

func main() {

	rootPath, err := os.Getwd()
	if err != nil {
		rootPath = "/tmp/"
	}
	cachePath := filepath.Join(rootPath, "cache-dir")
	simplecache.SetConfig(simplecache.Config{StorePath: cachePath})

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	database.DB = database.New(database.Postgres, psqlconn)
	defer database.DB.Close()

	// Run the command only if there is input arguments
	if len(os.Args) > 1 {
		cmd.Execute()
		return
	}

	// Otherwise start the API server
	api.ListenAndServeHTTP(os.Getenv("SERVING_ADDR"))
}
