package validators

import (
	"fmt"

	"github.com/archway-network/valuter/database"
)

func (v *ValidatorRecord) GetFirstSignedBlockHeight() (uint64, error) {

	SQL := fmt.Sprintf(`
		SELECT 
			MIN("%s") AS "result"
		FROM "%s" 
		WHERE "%s" = $1`,
		database.FIELD_BLOCK_SIGNERS_BLOCK_HEIGHT,
		database.TABLE_BLOCK_SIGNERS,
		database.FIELD_BLOCK_SIGNERS_VAL_CONS_ADDR,
	)

	rows, err := database.DB.Query(SQL, database.QueryParams{v.ConsAddr})
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

func (v *ValidatorRecord) GetLatestSignedBlockHeight() (uint64, error) {

	SQL := fmt.Sprintf(`
		SELECT 
			MAX("%s") AS "result"
		FROM "%s" 
		WHERE "%s" = $1`,
		database.FIELD_BLOCK_SIGNERS_BLOCK_HEIGHT,
		database.TABLE_BLOCK_SIGNERS,
		database.FIELD_BLOCK_SIGNERS_VAL_CONS_ADDR,
	)

	rows, err := database.DB.Query(SQL, database.QueryParams{v.ConsAddr})
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

func (v *ValidatorRecord) GetTotalSignedBlocks() (uint64, error) {

	SQL := fmt.Sprintf(`
		SELECT 
			COUNT(*) AS "total"
		FROM "%s" 
		WHERE "%s" = $1`,
		database.TABLE_BLOCK_SIGNERS,
		database.FIELD_BLOCK_SIGNERS_VAL_CONS_ADDR,
	)

	rows, err := database.DB.Query(SQL, database.QueryParams{v.ConsAddr})
	if err != nil {
		return 0, err
	}
	if len(rows) == 0 ||
		rows[0] == nil ||
		rows[0]["result"] == nil {
		return 0, nil
	}

	return uint64(rows[0]["total"].(int64)), nil
}
