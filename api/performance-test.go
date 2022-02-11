package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/archway-network/valuter/tasks"
	"github.com/archway-network/valuter/tools"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*
* This function implements GET /challenges/uptime
 */
func GetPerformanceTestWinners(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	winnersList, err := tasks.GetPerformanceTestWinners()

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, winnersList.GetItems())
}

/*-------------*/
/*
* This function implements GET /challenges/uptime/:burst_index
 */
func GetPerformanceTestWinnersPerLoadBurst(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	burstIndex, err := strconv.Atoi(params.ByName("burst_index"))
	if err != nil {
		burstIndex = 0
	}
	winnersList, err := tasks.GetPerformanceTestWinnersPerLoadBurst(burstIndex)

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, winnersList.GetItems())
}

/*-------------*/
