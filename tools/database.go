package tools

import (
	"log"
	"net/http"
	"strconv"

	"github.com/archway-network/valuter/configs"
	"github.com/archway-network/valuter/types"
)

/*------------------------------*/

func GetLimitOffsetFromHttpReq(req *http.Request) types.DBLimitOffset {
	qryParams := req.URL.Query()

	page := 1
	if _, ok := qryParams["page"]; ok {

		var err error
		page, err = strconv.Atoi(qryParams["page"][0])
		if err != nil {
			log.Printf("Error in page number: %v", err)
			page = 1
		}
		if page <= 0 {
			page = 1
		}
	}

	offset := (uint64(page) - 1) * configs.Configs.API.RowsPerPage

	return types.DBLimitOffset{
		Limit:  configs.Configs.API.RowsPerPage,
		Offset: offset,
		Page:   uint64(page),
	}
}

/*------------------------------*/
