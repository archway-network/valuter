package validators

import (
	"fmt"
	"sort"

	"github.com/archway-network/cosmologger/database"
	cosmoLogTx "github.com/archway-network/cosmologger/tx"
	"github.com/archway-network/valuter/blocks"
	"github.com/archway-network/valuter/tools"
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
	// Calculating uptime
	// Total blocks is more accurate as `cosmologger` might miss some blocks under some cirscumstance
	latestBlockHeight, err := blocks.GetLatestBlockHeight()
	if err != nil {
		return ValidatorInfo{}, err
	}
	return v.GetValidatorInfoByBlockHeightRange(1, latestBlockHeight)

}

func (v *ValidatorRecord) GetValidatorInfoByBlockHeightRange(beginHeight, endHeight uint64) (ValidatorInfo, error) {
	var vInfo ValidatorInfo
	var err error

	vInfo.ConsAddr = v.ConsAddr
	vInfo.OprAddr = v.OprAddr
	vInfo.AccAddr = v.AccAddr
	vInfo.Moniker = v.Moniker

	vInfo.FirstSignedBlockHeight, err = v.GetFirstSignedBlockHeightWithBegin(beginHeight)
	if err != nil {
		return vInfo, err
	}

	vInfo.TotalSignedBlocks, err = v.GetTotalSignedBlocksWithHeightRange(beginHeight, endHeight)
	if err != nil {
		return vInfo, err
	}

	// Calculating uptime
	totalBlocks := endHeight - beginHeight

	expectedSignedBlocks := totalBlocks - vInfo.FirstSignedBlockHeight
	if expectedSignedBlocks == 0 {
		vInfo.UpTime = 0
	} else {
		vInfo.UpTime = float32(vInfo.TotalSignedBlocks) / float32(expectedSignedBlocks)
	}
	vInfo.UpTimeAll = float32(vInfo.TotalSignedBlocks) / float32(totalBlocks)

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
	pagination := tools.GetPagination(totalRows, limitOffset.Page)

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
		cosmoLogTx.ACTION_UNJAIL,

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

func GetJoinedAfterGenesisValidators() ([]ValidatorRecord, error) {

	var validatorsList []ValidatorRecord

	txs, err := tx.GetAllTxsByAction(cosmoLogTx.ACTION_CREATE_VALIDATOR)
	if err != nil {
		return validatorsList, err
	}
	for i := range txs {

		v, err := GetValidatorByOprAddr(txs[i].Validator)
		if err != nil {
			return validatorsList, err
		}
		validatorsList = append(validatorsList, v)
	}

	return validatorsList, nil
}

func GetAllValidators() ([]ValidatorRecord, error) {

	rows, err := database.DB.Load(database.TABLE_VALIDATORS, nil)
	if err != nil {
		return nil, err
	}

	validators := DBRowToValidatorRecords(rows)
	return validators, err
}

func GetAllValidatorsWithInfoByBlockHeightRange(beginHeight, endHeight uint64) ([]ValidatorInfo, error) {

	var valInfoList []ValidatorInfo
	listOfValidators, err := GetAllValidators()
	if err != nil {
		return nil, err
	}
	for i := range listOfValidators {

		// fmt.Printf("Validator #%5d \tMoniker: `%50s` ...\t", i, listOfValidators[i].Moniker)
		valInfo, err := listOfValidators[i].GetValidatorInfoByBlockHeightRange(beginHeight, endHeight)
		// fmt.Printf("Done\n")

		if err != nil {
			return valInfoList, err
		}
		valInfoList = append(valInfoList, valInfo)
	}

	sort.Slice(valInfoList, func(i, j int) bool {
		return valInfoList[i].UpTime > valInfoList[j].UpTime
	})

	return valInfoList, nil
}

func GetAllValidatorsWithInfo() ([]ValidatorInfo, error) {

	latestBlockHeight, err := blocks.GetLatestBlockHeight()
	if err != nil {
		return []ValidatorInfo{}, err
	}

	return GetAllValidatorsWithInfoByBlockHeightRange(1, latestBlockHeight)
}
