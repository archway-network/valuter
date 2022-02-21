package blocksigners

import (
	"fmt"

	"github.com/archway-network/cosmologger/database"
	"github.com/archway-network/valuter/validators"
)

func GetSignersByBlockHeight(height uint64) ([]validators.ValidatorRecord, error) {

	SQL := fmt.Sprintf(`
		SELECT 
			v.* 
		FROM 
			"%s"	AS	b,
			"%s"	AS	v
		WHERE 
			b."%s" = $1 AND
			b."%s" = v."%s"`,

		database.TABLE_BLOCK_SIGNERS,
		database.TABLE_VALIDATORS,

		database.FIELD_BLOCK_SIGNERS_BLOCK_HEIGHT,

		database.FIELD_BLOCK_SIGNERS_VAL_CONS_ADDR,
		database.FIELD_VALIDATORS_CONS_ADDR,
	)

	rows, err := database.DB.Query(SQL, database.QueryParams{height})

	if err != nil {
		return nil, err
	}

	return validators.DBRowToValidatorRecords(rows), nil
}
