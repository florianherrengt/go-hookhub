package routers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/florianherrengt/hubhook/config"
	"github.com/florianherrengt/hubhook/pkg/datasource"
	"github.com/florianherrengt/hubhook/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// NewRouter create a gin router
func NewRouter() *gin.Engine {
	r := gin.New()

	r.GET("/events", func(ctx *gin.Context) {
		var hookEvents []models.HookEvent
		result := datasource.DB.Find(&hookEvents)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"items": hookEvents})
	})
	r.POST("/hook", func(ctx *gin.Context) {
		contentType := ctx.Request.Header.Get("content-type")
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		hookEvent := models.HookEvent{ID: uuid.New(), Body: string(body), ContentType: contentType}
		hookEventMarshal, err := json.Marshal(hookEvent)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		err = datasource.StreamClient.Publish(config.Config.PubSub.EventName.NewIncomingRequest, []byte(string(hookEventMarshal)))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		ctx.JSON(http.StatusOK, gin.H{"ok": 1})
	})
	return r
}
