package blocks

import (
	"fmt"

	"github.com/archway-network/cosmologger/database"
	"github.com/archway-network/valuter/types"
)

func GetLatestBlock() (types.BlockRecord, error) {

	SQL := fmt.Sprintf(`
		SELECT 
			*
		FROM "%s" 
		ORDER BY "%s" DESC
		LIMIT 1`,
		database.TABLE_BLOCKS,
		database.FIELD_BLOCKS_HEIGHT,
	)

	rows, err := database.DB.Query(SQL, database.QueryParams{})
	if err != nil {
		return types.BlockRecord{}, err
	}

	if len(rows) == 0 {
		return types.BlockRecord{}, nil
	}

	return DBRowToBlockRecord(rows[0]), nil
}

func GetLatestBlockHeight() (uint64, error) {

	SQL := fmt.Sprintf(
		`SELECT MAX("%s") AS "result" FROM "%s"`,

		database.FIELD_BLOCKS_HEIGHT,
		database.TABLE_BLOCKS,
	)

	rows, err := database.DB.Query(SQL, database.QueryParams{})
	if err != nil {
		return 0, err
	}

	if len(rows) == 0 ||
		rows[0] == nil ||
		rows[0]["result"] == nil {
		return 0, nil
	}

	return uint64(rows[0]["result"].(int64)), nil
}

func GetBlockByHeight(height uint64) (types.BlockRecord, error) {

	SQL := fmt.Sprintf(`
		SELECT 
			*
		FROM "%s" 
		WHERE "%s" = $1`,
		database.TABLE_BLOCKS,
		database.FIELD_BLOCKS_HEIGHT,
	)

	rows, err := database.DB.Query(SQL, database.QueryParams{height})
	if err != nil {
		return types.BlockRecord{}, err
	}

	if len(rows) == 0 {
		return types.BlockRecord{}, nil
	}

	return DBRowToBlockRecord(rows[0]), nil
}

func GetTotalBlocks() (uint64, error) {

	SQL := fmt.Sprintf(`
		SELECT 
			COUNT(*) AS "total"
		FROM "%s"`,
		database.TABLE_BLOCKS,
	)

	rows, err := database.DB.Query(SQL, database.QueryParams{})
	if err != nil {
		return 0, err
	}
	if len(rows) == 0 ||
		rows[0] == nil ||
		rows[0]["total"] == nil {
		return 0, nil
	}

	return uint64(rows[0]["total"].(int64)), nil
}
