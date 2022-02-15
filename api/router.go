package api

import (
	"log"
	"net/http"

	routing "github.com/julienschmidt/httprouter"
)

/*-------------------------*/

func setupRouter() *routing.Router {

	var router = routing.New()

	router.GET("/", IndexPage)
	router.GET("/configs", GetAllConfigs)
	// router.GET("/ui/*file_path", UI)

	// router.GET("/winners", CheckAPIKey(GetWinners))
	router.GET("/winners", GetWinners)
	router.GET("/winners/:address", GetWinner)

	router.GET("/challenges", GetListOfChallenges)
	router.GET("/challenges/gov", GetGovWinners)
	router.GET("/challenges/gov/:proposal_id", GetGovWinnersPerProposal)
	router.GET("/challenges/staking", GetStakingWinners)
	router.GET("/challenges/node-upgrade", GetNodeUpgradeWinners)
	router.GET("/challenges/validators-genesis", GetNodeUpgradeWinners)
	router.GET("/challenges/validators-joined", GetNodeUpgradeWinners)
	router.GET("/challenges/jail-unjail", GetNodeUpgradeWinners)
	router.GET("/challenges/uptime", GetPerformanceTestWinners)
	router.GET("/challenges/uptime/:burst_index", GetPerformanceTestWinnersPerLoadBurst)

	router.GET("/validators", GetValidators)
	router.GET("/validators/validator/:address", GetValidator)
	router.GET("/validators/genesis", GetGenesisValidators)
	router.GET("/validators/joined", GetJoinedAfterGenesisValidators)
	router.GET("/validators/unjailed", GetUnjailedValidators)

	return router
}

/*-------------------------*/

// ListenAndServeHTTP serves the APIs and the ui
func ListenAndServeHTTP(addr string) {

	router := setupRouter()
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("[INFO ] Serving on %s", addr)

	log.Fatal(http.ListenAndServe(addr, router))
}

/*-------------------------*/
