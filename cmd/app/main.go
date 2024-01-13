package main

import (
	"github.com/forg002-ctrl/managment_system/pkg/config"
	"github.com/forg002-ctrl/managment_system/pkg/router"
)

func main() {
	cfg := config.Init();
	
	r := router.SetupRouter()

	r.Run(":" + cfg.Port)
}