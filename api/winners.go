package api

import (
	"net/http"

	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*
* This function implements GET /winners
 */
func GetWinners(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	// limitOffset := tools.GetLimitOffset(req)

	/*------*/

	// totalRows := int64(0)
	// {
	// 	SQL := `SELECT COUNT(*) AS "total" FROM "scooters"`
	// 	rows, err := global.DB.Query(SQL, database.QueryParams{})
	// 	if err != nil {
	// 		log.Printf("Error in db query: %v", err)
	// 		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	totalRows = rows[0]["total"].(int64)
	// }

	// totalPages := int64(math.Ceil(float64(totalRows) / float64(global.RowsPerPage)))
	// pagination := map[string]interface{}{
	// 	"current_page":  limitOffset.Page,
	// 	"total_pages":   totalPages,
	// 	"total_entries": totalRows,
	// }

	// /*------*/

	// SQL := `SELECT *
	// 		FROM
	// 			"scooters"
	// 		LIMIT $1 OFFSET $2`

	// rows, err := global.DB.Query(SQL, database.QueryParams{limitOffset.Limit, limitOffset.Offset})
	// if err != nil {
	// 	log.Printf("Error in db query: %v", err)
	// 	http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// tools.SendJSON(resp, map[string]interface{}{"pagination": pagination, "rows": rows})
}

/*-------------*/

/*
* This function implements GET /winners/:uuid
 */
func GetWinner(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	// addr := params.ByName("addr")

	// /*------*/

	// SQL := `SELECT * FROM "scooters" WHERE "uuid" = $1`

	// rows, err := global.DB.Query(SQL, database.QueryParams{uuid})
	// if err != nil {
	// 	log.Printf("Error in db query: %v", err)
	// 	http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// if rows == nil || len(rows) == 0 {
	// 	http.Error(resp, "Scooter not found!", http.StatusNotFound)
	// 	return
	// }

	// tools.SendJSON(resp, rows[0])
}

/*-------------*/
