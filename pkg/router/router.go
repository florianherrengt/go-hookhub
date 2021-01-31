package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/florianherrengt/hubhook/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// NewRouter create a gin router
func NewRouter() *gin.Engine {
	r := gin.New()

	r.POST("/hook", func(ctx *gin.Context) {
		hookEvent := models.HookEvent{ID: uuid.New(), Payload: "hello"}
		hookEventString, err := json.Marshal(hookEvent)
		fmt.Println(string(hookEventString), err)
		// datasource.StreamClient.Publish(config.Config.PubSub.EventName.NewIncomingRequest, []byte("new request came in"))
		ctx.JSON(http.StatusOK, gin.H{"ok": 1})
	})
	return r
}
