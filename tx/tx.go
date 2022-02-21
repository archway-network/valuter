package tx

import (
	"fmt"

	"github.com/archway-network/cosmologger/database"
	"github.com/archway-network/valuter/tools"
	"github.com/archway-network/valuter/types"
)

// With Pagination
func GetTxsByAction(
	action string,
	limitOffset types.DBLimitOffset) ([]types.TxRecord, types.Pagination, error) {
	return getTxsWithPagination(
		database.RowType{
			database.FIELD_TX_EVENTS_ACTION: action,
		},
		limitOffset)
}

// Without Pagination
func GetAllTxsByAction(action string) ([]types.TxRecord, error) {
	return getTxs(
		database.RowType{
			database.FIELD_TX_EVENTS_ACTION: action,
		})
}

func getTxsWithPagination(
	conditions database.RowType,
	limitOffset types.DBLimitOffset) ([]types.TxRecord, types.Pagination, error) {

	SQL := fmt.Sprintf(`SELECT * FROM "%s" WHERE 1 = 1 `, database.TABLE_TX_EVENTS)

	var params database.QueryParams
	paramCounter := 1
	for fieldName, value := range conditions {
		SQL += fmt.Sprintf(` AND "%s" = $%d `, fieldName, paramCounter)
		paramCounter++
		params = append(params, value)
	}

	// Prepare pagination
	totalRows := uint64(0)
	{
		pSQL := fmt.Sprintf(`SELECT COUNT(*) AS "total" FROM (%s) AS tmp`, SQL)
		rows, err := database.DB.Query(pSQL, params)
		if err != nil {
			return nil, types.Pagination{}, err
		}
		totalRows = uint64(rows[0]["total"].(int64))
	}

	pagination := tools.GetPagination(totalRows, limitOffset.Page)

	SQL += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, paramCounter, paramCounter+1)
	params = append(params, limitOffset.Limit)
	params = append(params, limitOffset.Offset)

	// Order by who is first
	SQL += fmt.Sprintf(` ORDER BY "%s" ASC `, database.FIELD_TX_EVENTS_HEIGHT)

	fmt.Println(SQL)

	rows, err := database.DB.Query(SQL, params)
	if err != nil {
		return nil, types.Pagination{}, err
	}

	return DBRowsToTxRecords(rows), pagination, err
}

func getTxs(conditions database.RowType) ([]types.TxRecord, error) {

	SQL := fmt.Sprintf(`SELECT * FROM "%s" WHERE 1 = 1 `, database.TABLE_TX_EVENTS)

	var params database.QueryParams
	paramCounter := 1
	for fieldName, value := range conditions {
		SQL += fmt.Sprintf(` AND "%s" = $%d `, fieldName, paramCounter)
		paramCounter++
		params = append(params, value)
	}

	// Order by who is first
	SQL += fmt.Sprintf(` ORDER BY "%s" ASC `, database.FIELD_TX_EVENTS_HEIGHT)

	rows, err := database.DB.Query(SQL, params)
	if err != nil {
		return nil, err
	}

	return DBRowsToTxRecords(rows), err
}
