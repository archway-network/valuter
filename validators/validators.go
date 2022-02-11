package validators

import (
	"fmt"
	"math"

	"github.com/archway-network/valuter/blocks"
	"github.com/archway-network/valuter/configs"
	"github.com/archway-network/valuter/database"
	"github.com/archway-network/valuter/tx"
	"github.com/archway-network/valuter/types"
)

func GetValidatorByConsAddr(valConsAddr string) (ValidatorRecord, error) {

	SQL := fmt.Sprintf(`
		SELECT *
		FROM "%s" 
		WHERE "%s" = $1`,
		database.TABLE_VALIDATORS,
		database.FIELD_VALIDATORS_CONS_ADDR,
	)

	rows, err := database.DB.Query(SQL, database.QueryParams{valConsAddr})
	if err != nil {
		return ValidatorRecord{}, err
	}
	if len(rows) == 0 {
		return ValidatorRecord{}, nil
	}

	return DBRowToValidatorRecord(rows[0]), nil
}

func GetValidatorByOprAddr(valOprAddr string) (ValidatorRecord, error) {

	SQL := fmt.Sprintf(`
		SELECT *
		FROM "%s" 
		WHERE "%s" = $1`,
		database.TABLE_VALIDATORS,
		database.FIELD_VALIDATORS_OPR_ADDR,
	)

	rows, err := database.DB.Query(SQL, database.QueryParams{valOprAddr})
	if err != nil {
		return ValidatorRecord{}, err
	}
	if len(rows) == 0 {
		return ValidatorRecord{}, nil
	}

	return DBRowToValidatorRecord(rows[0]), nil
}

func (v *ValidatorRecord) GetValidatorInfo() (ValidatorInfo, error) {
	var vInfo ValidatorInfo
	var err error

	vInfo.ConsAddr = v.ConsAddr
	vInfo.OprAddr = v.OprAddr

	vInfo.FirstSignedBlockHeight, err = v.GetFirstSignedBlockHeight()
	if err != nil {
		return vInfo, err
	}

	vInfo.TotalSignedBlocks, err = v.GetTotalSignedBlocks()
	if err != nil {
		return vInfo, err
	}

	// Calculating uptime
	// Total blocks is more accurate as `cosmologger` might miss some blocks under some cirscumstance
	totalLoggedBlocks, err := blocks.GetTotalBlocks()
	if err != nil {
		return vInfo, err
	}
	expectedSignedBlocks := totalLoggedBlocks - vInfo.FirstSignedBlockHeight
	vInfo.UpTime = float32(vInfo.TotalSignedBlocks) / float32(expectedSignedBlocks)

	return vInfo, nil

}

func GetValidatorsWithPagination(limitOffset types.DBLimitOffset) ([]ValidatorRecord, types.Pagination, error) {

	// Prepare pagination
	totalRows := uint64(0)
	{
		SQL := fmt.Sprintf(`SELECT COUNT(*) AS "total" FROM "%s"`,
			database.TABLE_VALIDATORS,
		)
		rows, err := database.DB.Query(SQL, database.QueryParams{})
		if err != nil {
			return nil, types.Pagination{}, err
		}
		totalRows = uint64(rows[0]["total"].(int64))
	}

	totalPages := uint64(math.Ceil(float64(totalRows) / float64(configs.Configs.API.RowsPerPage)))
	pagination := types.Pagination{
		CurrentPage: limitOffset.Page,
		TotalPages:  totalPages,
		TotalRows:   totalRows,
	}

	/*------*/

	SQL := fmt.Sprintf(`SELECT * FROM "%s" LIMIT $1 OFFSET $2`, database.TABLE_VALIDATORS)

	rows, err := database.DB.Query(SQL,
		database.QueryParams{
			limitOffset.Limit,
			limitOffset.Offset,
		})
	if err != nil {
		return nil, types.Pagination{}, err
	}

	validators := DBRowToValidatorRecords(rows)
	return validators, pagination, err
}

func GetUnjailedValidators() ([]ValidatorWithTx, error) {

	SQL := fmt.Sprintf(`
			SELECT * FROM (
				SELECT DISTINCT ON(v."%s") * 
				FROM 
					"%s" AS v,
					"%s" AS t
				WHERE 
					t."%s" = '%s' AND
					v."%s" = t."%s"
				ORDER BY v."%s"
			) AS "tmp"
			ORDER BY "%s" ASC`,

		database.FIELD_VALIDATORS_OPR_ADDR,

		database.TABLE_VALIDATORS,
		database.TABLE_TX_EVENTS,

		database.FIELD_TX_EVENTS_ACTION,
		tx.ACTION_UNJAIL,

		database.FIELD_VALIDATORS_OPR_ADDR,
		database.FIELD_TX_EVENTS_SENDER, // The ValOper address is set to the `sender` field for `unjail`

		database.FIELD_VALIDATORS_OPR_ADDR,

		database.FIELD_TX_EVENTS_HEIGHT,
	)

	rows, err := database.DB.Query(SQL, database.QueryParams{})
	if err != nil {
		return nil, err
	}

	return DBRowToValidatorWithTxs(rows), err
}
