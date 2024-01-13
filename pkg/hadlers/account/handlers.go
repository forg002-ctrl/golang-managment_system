package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/forg002-ctrl/managment_system/pkg/models"
)

func CreateAccount(c *gin.Context) {
	c.JSON(http.StatusOK, models.Account{
		ID: 0,
		Username: "test",
	})
}