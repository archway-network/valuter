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
	router.GET("/ui/*file_path", UI)

	// router.GET("/winners", CheckAPIKey(GetWinners))
	router.GET("/winners", GetWinners)
	router.GET("/winners/:address", GetWinner)

	router.GET("/validators", GetValidators)
	router.GET("/validators/validator/:address", GetValidator)
	router.GET("/validators/genesis", GetGenesisValidators)
	router.GET("/validators/joined", GetJoinedLaterValidators)

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
