package main

import (
	"fmt"
	"log"

	"github.com/forg002-ctrl/managment_system/pkg/config"
	"github.com/forg002-ctrl/managment_system/pkg/db"
	"github.com/forg002-ctrl/managment_system/pkg/router"
	"github.com/forg002-ctrl/managment_system/pkg/storage"
)

func main() {
	cfg := config.Init();

	connectionOptions := &storage.DBConnectionOptions{
		Host: cfg.DBhost,
		User: cfg.DBuser,
		Password: cfg.DBpassword,
		Name: cfg.DBname,
		Port: cfg.DBport,
	}
	storageClient := &storage.PostgresClient{}
	storageClient.Init(connectionOptions)

	err := db.InitTables(storageClient)
	if err != nil {
		log.Fatal(err)
	}

	r := router.SetupRouter()

	port := cfg.Port
	if port == "" {
		port = "3000"
	}
	addr := fmt.Sprintf(":%s", port)
 
	r.Run(addr)
}