package api

import (
	"log"
	"net/http"

	"github.com/archway-network/valuter/tasks"
	"github.com/archway-network/valuter/tools"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*
* This function implements GET /challenges/staking
 */
func GetGovWinners(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	winnersList, err := tasks.GetGovWinners()

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, winnersList.GetItems())
}

/*-------------*/
