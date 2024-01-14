package main

import (
	"fmt"

	"github.com/forg002-ctrl/managment_system/pkg/config"
	"github.com/forg002-ctrl/managment_system/pkg/router"
)

func main() {
	cfg := config.Init();

	r := router.SetupRouter()

	port := cfg.Port
	addr := fmt.Sprintf(":%s", port)
 
	r.Run(addr)
}