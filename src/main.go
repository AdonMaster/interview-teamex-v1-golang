package main

import (
	"context"
	"interview-teamex-v1/src/config"
	"interview-teamex-v1/src/db"
	"interview-teamex-v1/src/mids"
	"interview-teamex-v1/src/repo"
	"interview-teamex-v1/src/router"
	"log"
	"net/http"
)

func main() {

	//
	ctx := context.Background()
	var err error

	// init environment
	config.Init()

	// db + migration + seed
	conn, err := db.Init(ctx)
	if err != nil {
		panic(err)
	}

	// init repository
	repository := repo.New(conn)

	// router
	routerInstance := mids.Cors(router.Init(&repository))

	// let's serve this
	log.Printf("\n\nServing on http://localhost:%s", config.Env.Port)
	log.Fatal(http.ListenAndServe(":"+config.Env.Port, routerInstance))

}
