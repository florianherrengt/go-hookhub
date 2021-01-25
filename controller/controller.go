package controller

import (
	"log"
	"net/http"

	"github.com/florianherrengt/gohubhook/models"
	"github.com/gin-gonic/gin"
)

type queryParams struct {
	Order string `form:"order" default:"desc"`
	Limit int    `form:"limit" binding:"max=100"`
}

type uriParams struct {
	ID string `uri:"id"`
}

// CreateEvent in the db
func CreateEvent(ctx *gin.Context) {
	newEvent := models.IncomingRequest{Payload: "helo "}
	result := models.DB.Create(&newEvent)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": newEvent})
}

// ListEvents from the db
func ListEvents(ctx *gin.Context) {
	var allEvents []models.IncomingRequest
	models.DB.Find(&allEvents)
	ctx.JSON(http.StatusOK, gin.H{"data": allEvents})
}

// ShowAccount wip
func ShowAccount(ctx *gin.Context) {
	var queryParams queryParams
	var uriParams uriParams
	if err := ctx.Bind(&queryParams); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.BindUri(&uriParams)

	log.Println(&queryParams, ctx)
	order := queryParams.Order
	if order == "" {
		order = "desc"
	}
	log.Println(uriParams)

	ctx.JSON(http.StatusOK, gin.H{"order": order})
}
