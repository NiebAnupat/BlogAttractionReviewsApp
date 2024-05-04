package main

import (
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/config"
	database "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/database"
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/server"
)

func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)
	sever := server.NewFiberServer(conf, db)
	sever.Start()
}
