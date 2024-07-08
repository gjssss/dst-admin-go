package router

import (
	"dst-admin-go/api"
	"github.com/gin-gonic/gin"
)

func initMsRouter(router *gin.RouterGroup) {

	msApi := api.MsApi{}
	ms := router.Group("")
	{
		ms.GET("/sync/level/config", msApi.SyncLevelConfig)
		ms.GET("/sync/slave/status", msApi.SyncSlaveLevelStatus)
		ms.GET("/sync/slave/start", msApi.SyncSlaveStart)
		ms.GET("/sync/slave/stop", msApi.SyncSlaveStop)
		ms.GET("/slave/status", msApi.GetSlaveLevelStatus)
	}

}
