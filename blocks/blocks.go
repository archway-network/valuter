package blocks

import (
	"fmt"

	"github.com/archway-network/valuter/database"
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
