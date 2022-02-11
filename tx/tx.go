package tx

import (
	"fmt"
	"math"

	"github.com/archway-network/valuter/configs"
	"github.com/archway-network/valuter/database"
	"github.com/archway-network/valuter/types"
)

func GetTxsByAction(
	action string,
	limitOffset types.DBLimitOffset) ([]types.TxRecord, types.Pagination, error) {
	return getTxsWithPagination(
		database.RowType{
			database.FIELD_TX_EVENTS_ACTION: action,
		},
		limitOffset)
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

	totalPages := uint64(math.Ceil(float64(totalRows) / float64(configs.Configs.API.RowsPerPage)))
	pagination := types.Pagination{
		CurrentPage: limitOffset.Page,
		TotalPages:  totalPages,
		TotalRows:   totalRows,
	}

	rows, err := database.DB.Query(SQL,
		database.QueryParams{
			limitOffset.Limit,
			limitOffset.Offset,
		})
	if err != nil {
		return nil, types.Pagination{}, err
	}

	return DBRowsToTxRecords(rows), pagination, err
}
