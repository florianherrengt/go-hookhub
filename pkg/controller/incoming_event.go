package controller

import (
	"net/http"

	"github.com/florianherrengt/hubhook/config"
	"github.com/florianherrengt/hubhook/pkg/datasource"
	"github.com/gin-gonic/gin"
)

// IncomingEventGetHandler handle GET requests from /hook
func IncomingEventGetHandler(ctx *gin.Context) {
	datasource.StreamClient.Publish(config.Config.PubSub.EventName.NewIncomingRequest, []byte("new request came in"))
	ctx.JSON(http.StatusOK, gin.H{"ok": 1})
}
