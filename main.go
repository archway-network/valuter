package main

import (
	"fmt"
	"os"

	"github.com/archway-network/valuter/api"
	"github.com/archway-network/valuter/configs"
	"github.com/archway-network/valuter/database"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

	// conn, err := Connect()
	// if err != nil {
	// 	log.Fatalf("Did not connect: %s", err)
	// }
	// defer conn.Close()

	/*-------------*/

	// SetBech32Prefixes()

	/*-------------*/

	// vals, pg, err := tasks.GetActiveValidators(types.DBLimitOffset{Limit: 50})
	// if err != nil {
	// 	panic(err)
	// }

	// for i := range vals {
	// 	fmt.Printf("Addr: %s \tTotal Signed: %d\n", vals[i].ConsAddr, vals[i].TotalSignedBlocks)
	// }

	// fmt.Printf("pg: %v\n", pg)

	fmt.Println("\nCiao bello!")

	api.ListenAndServeHTTP(os.Getenv("SERVING_ADDR"))

}

// func Connect() (*grpc.ClientConn, error) {

// 	if configs.Configs.GRPC.TLS {
// 		creds := credentials.NewTLS(&tls.Config{})
// 		return grpc.Dial(configs.Configs.GRPC.Server, grpc.WithTransportCredentials(creds))
// 	}
// 	return grpc.Dial(configs.Configs.GRPC.Server, grpc.WithInsecure())

// }

func SetBech32Prefixes() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(configs.Configs.Bech32Prefix.Account.Address, configs.Configs.Bech32Prefix.Account.PubKey)
	config.SetBech32PrefixForValidator(configs.Configs.Bech32Prefix.Validator.Address, configs.Configs.Bech32Prefix.Validator.PubKey)
	config.SetBech32PrefixForConsensusNode(configs.Configs.Bech32Prefix.Consensus.Address, configs.Configs.Bech32Prefix.Consensus.PubKey)
	config.Seal()
}
