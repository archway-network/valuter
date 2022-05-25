package contracts

// Only contractRewardsAmount will be sufficient for: https://philabs.notion.site/Build-a-dApp-to-Earn-Max-Network-Rewards-94152e5648574df4b1462d15d1d6cea3

// https://philabs.notion.site/Build-a-dApp-To-Subsidize-Users-Fees-1674bce247734b83a377cfb920e7965b
// if gasRebateToUser == true:
// MAX(gasConsumed)

import (
	"fmt"

	"github.com/archway-network/cosmologger/block"
	"github.com/archway-network/cosmologger/database"
)

func GetMaxNetworkRewardsTopContracts(top, beginHeight, endHeight uint64) ([]block.ContractRecord, error) {

	SQL := fmt.Sprintf(`
		SELECT * FROM (
			SELECT DISTINCT ON ("%s") *
			FROM "%s" 
			WHERE 
				"%s" >= $1 AND 
				"%s" <= $2
			LIMIT $3 ) AS "tmp"
			ORDER BY "%s" DESC`,

		database.FIELD_CONTRACTS_CONTRACT_ADDRESS,
		database.TABLE_CONTRACTS,

		database.FIELD_CONTRACTS_BLOCK_HEIGHT,
		database.FIELD_CONTRACTS_BLOCK_HEIGHT,

		database.FIELD_CONTRACTS_CONTRACT_REWARDS_AMOUNT,
	)

	params := database.QueryParams{beginHeight, endHeight, top}
	rows, err := database.DB.Query(SQL, params)
	if err != nil {
		return []block.ContractRecord{}, err
	}

	return DBRowToContractRecords(rows), nil
}

func GetSubsidizeUsersFeesTopContracts(top, beginHeight, endHeight uint64) ([]block.ContractRecord, error) {

	SQL := fmt.Sprintf(`
		SELECT * FROM (
			SELECT DISTINCT ON ("%s") *
			FROM "%s" 
			WHERE
				"%s" = true AND 
				"%s" >= $1 AND 
				"%s" <= $2
			LIMIT $3 ) AS "tmp"
			ORDER BY "%s" DESC`,

		database.FIELD_CONTRACTS_CONTRACT_ADDRESS,
		database.TABLE_CONTRACTS,

		database.FIELD_CONTRACTS_GAS_REBATE_TO_USER,
		database.FIELD_CONTRACTS_BLOCK_HEIGHT,
		database.FIELD_CONTRACTS_BLOCK_HEIGHT,

		database.FIELD_CONTRACTS_GAS_CONSUMED,
	)

	params := database.QueryParams{beginHeight, endHeight, top}
	rows, err := database.DB.Query(SQL, params)
	if err != nil {
		return []block.ContractRecord{}, err
	}

	return DBRowToContractRecords(rows), nil
}
