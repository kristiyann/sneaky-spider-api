package main

import (
	"log"
	"os"

	"github.com/kristiyann/af1-spider-web-app/alerts"
	"github.com/kristiyann/af1-spider-web-app/api"
	"github.com/kristiyann/af1-spider-web-app/backend"
	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/httpclient"
	"github.com/kristiyann/af1-spider-web-app/logic"
	"github.com/kristiyann/af1-spider-web-app/util"
	"github.com/nedpals/supabase-go"
)

func main() {
	dbConn := util.LoadEnvVar(constants.EnvPostgresUrl)

	dbLogger := log.New(os.Stdout, "[sql] ", log.LstdFlags|log.Llongfile)
	querier := backend.NewQuerier(dbConn, dbLogger)
	db := backend.NewPostgresDb(*querier, dbConn)

	nikeClient := httpclient.New(constants.NikeBaseUrl)
	snkrsClient := httpclient.New(constants.NikeBaseUrl)

	supabaseUrl := util.LoadEnvVar(constants.EnvSupabaseApiUrl)
	supabaseKey := util.LoadEnvVar(constants.EnvSupabaseApiKey)

	supabaseClient := supabase.CreateClient(supabaseUrl, supabaseKey)

	logic := logic.New(db, nikeClient, snkrsClient)

	serverLogger := log.New(os.Stdout, "[API] ", log.LstdFlags|log.Llongfile)
	server := api.NewServer(util.LoadEnvVar(constants.EnvApiAddr), logic, supabaseClient, db, serverLogger)

	go alerts.Init()
	server.Run()
}
