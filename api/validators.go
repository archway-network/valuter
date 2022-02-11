package api

import (
	"log"
	"net/http"

	"github.com/archway-network/valuter/tools"
	"github.com/archway-network/valuter/validators"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*
* This function implements GET /validators
 */
func GetValidators(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	limitOffset := tools.GetLimitOffsetFromHttpReq(req)

	validators, pagination, err := validators.GetValidatorsWithPagination(limitOffset)

	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp,
		map[string]interface{}{
			"pagination": pagination,
			"rows":       validators,
		})
}

/*-------------*/
/*
* This function implements GET /validator/:address
 */
func GetValidator(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	address := params.ByName("address")
	var val validators.ValidatorRecord
	var err error

	if validators.IsConsAddr(address) {
		val, err = validators.GetValidatorByConsAddr(address)

	} else if validators.IsOprAddr(address) {

		val, err = validators.GetValidatorByOprAddr(address)
	}
	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	valInfo, err := val.GetValidatorInfo()
	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, valInfo)
}

/*-------------*/

/*
* This function implements GET /validators/genesis
 */
func GetGenesisValidators(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	// limitOffset := tools.GetLimitOffsetFromHttpReq(req)

	// validators, pagination, err := tasks.GetGenesisValidators()

	// if err != nil {
	// 	log.Printf("Error in db query: %v", err)
	// 	http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// tools.SendJSON(resp,
	// 	map[string]interface{}{
	// 		"pagination": pagination,
	// 		"rows":       validators,
	// 	})
}

/*-------------*/

/*
* This function implements GET /validators/joined
 */
func GetJoinedLaterValidators(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	// limitOffset := tools.GetLimitOffsetFromHttpReq(req)

	// validators, pagination, err := tasks.GetJoinedAfterGenesisValidators()

	// if err != nil {
	// 	log.Printf("Error in db query: %v", err)
	// 	http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// tools.SendJSON(resp,
	// 	map[string]interface{}{
	// 		"pagination": pagination,
	// 		"rows":       validators,
	// 	})
}

/*-------------*/

/*
* This function implements GET /validators/unjailed
 */
func GetUnjailedValidators(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	validators, err := validators.GetUnjailedValidators()

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, validators)
}

/*-------------*/
