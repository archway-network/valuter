package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/archway-network/valuter/blocksigners"
	"github.com/archway-network/valuter/participants"
	"github.com/archway-network/valuter/tasks"
	"github.com/archway-network/valuter/tools"
	"github.com/archway-network/valuter/validators"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*
* This function implements GET /validators
*                          GET /validators?beginHeight=xx&endHeight=xx
 */
func GetValidators(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	var valInfos []validators.ValidatorInfo
	var err error

	qryParams := req.URL.Query()
	if _, ok := qryParams["endHeight"]; ok {
		var beginHeight, endHeight uint64

		endHeight, err = strconv.ParseUint(qryParams["endHeight"][0], 10, 64)
		if err != nil {
			http.Error(resp, "Wrong endHeight: "+err.Error(), http.StatusBadRequest)
			return
		}

		if _, ok := qryParams["beginHeight"]; ok {
			beginHeight, err = strconv.ParseUint(qryParams["beginHeight"][0], 10, 64)
			if err != nil {
				http.Error(resp, "Wrong beginHeight: "+err.Error(), http.StatusBadRequest)
				return
			}
		}

		valInfos, err = validators.GetAllValidatorsWithInfoByBlockHeightRange(beginHeight, endHeight)

	} else {

		valInfos, err = validators.GetAllValidatorsWithInfo()
	}

	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	type ValOutput struct {
		validators.ValidatorInfo
		participants.ParticipantRecord
	}

	var output []ValOutput
	for i := range valInfos {
		pRec, _ := participants.GetParticipantByAddress(valInfos[i].AccAddr)
		output = append(output,
			ValOutput{
				ValidatorInfo:     valInfos[i],
				ParticipantRecord: pRec,
			})
	}

	tools.SendJSON(resp, output)

	// limitOffset := tools.GetLimitOffsetFromHttpReq(req)
	// validators, pagination, err := validators.GetValidatorsWithPagination(limitOffset)
}

/*-------------*/
/*
* This function implements GET /validators/validator/:address
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

	if valInfo.ConsAddr == "" {
		http.Error(resp, "Validator Not Found", http.StatusNotFound)
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
	listOfValidators, err := blocksigners.GetSignersByBlockHeight(1)

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, listOfValidators)
}

/*-------------*/

/*
* This function implements GET /validators/joined
 */
func GetJoinedAfterGenesisValidators(resp http.ResponseWriter, req *http.Request, params routing.Params) {
	// limitOffset := tools.GetLimitOffsetFromHttpReq(req)

	listOfValidators, err := validators.GetJoinedAfterGenesisValidators()

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, listOfValidators)
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

/*
* This function implements GET /challenges/validators-genesis
 */
func GetGenesisValidatorsWinners(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	winnersList, err := tasks.GetGenesisValidatorsWinners()

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, winnersList.GetItems())
}

/*-------------*/
/*
* This function implements GET /challenges/validators-joined
 */
func GetJoinedAfterGenesisValidatorsWinners(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	winnersList, err := tasks.GetJoinedAfterGenesisValidatorsWinners()

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, winnersList.GetItems())
}

/*-------------*/

/*
* This function implements GET /challenges/validators-joined
 */
func GetUnjailedValidatorsWinners(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	winnersList, err := tasks.GetUnjailedValidatorsWinners()

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, winnersList.GetItems())
}

/*-------------*/
