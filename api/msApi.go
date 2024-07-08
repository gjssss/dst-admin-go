package api

import (
	"dst-admin-go/ms"
	"dst-admin-go/utils/dstConfigUtils"
	"dst-admin-go/vo"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type MsApi struct {
}

func (c *MsApi) SyncLevelConfig(ctx *gin.Context) {
	config := dstConfigUtils.GetDstConfig()
	clusterName := config.Cluster
	clusterToken := homeService.GetClusterToken(clusterName)
	clusterIni := homeService.GetClusterIni(clusterName)
	modoverrides := homeService.GetMasterWorld(clusterName).Modoverrides
	syncLevelData := ms.SyncLevelData{
		Modoverrides: modoverrides,
		ClusterToken: clusterToken,
		MasterPort:   clusterIni.MasterPort,
		MasterIp:     clusterIni.MasterIp,
	}
	bytes, err := json.Marshal(syncLevelData)
	if err != nil {
		log.Panicln(err)
	}
	clients := ms.Server.GetClient()
	for i := range clients {
		ms.Server.SyncLevelToClient(clients[i], string(bytes))
	}
	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "success",
		Data: nil,
	})
}

func (c *MsApi) SyncSlaveStart(ctx *gin.Context) {
	clientName := ctx.Query("slave")
	if clientName == "" {
		clients := ms.Server.GetClient()
		for i := range clients {
			ms.Server.SyncStartToClient(clients[i], "start")
		}
	} else {
		ms.Server.SyncStartToClient(clientName, "start")
	}

	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "success",
		Data: nil,
	})
}

func (c *MsApi) SyncSlaveStop(ctx *gin.Context) {
	clientName := ctx.Query("slave")
	if clientName == "" {
		clients := ms.Server.GetClient()
		for i := range clients {
			ms.Server.SyncStartToClient(clients[i], "stop")
		}
	} else {
		ms.Server.SyncStartToClient(clientName, "stop")
	}

	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "success",
		Data: nil,
	})
}

type SlaveInfo struct {
	Name  string `json:"name"`
	Extra string `json:"extra"`
}

func (c *MsApi) GetSlaveLevelStatus(ctx *gin.Context) {
	clients := ms.Server.GetClient()
	clientsExtra := ms.Server.GetClientExtra()
	clientInfoList := make([]SlaveInfo, 0)
	for i := range clients {
		name := clients[i]
		if value, ok := clientsExtra[name]; ok {
			clientInfo := SlaveInfo{
				Name:  name,
				Extra: value,
			}
			clientInfoList = append(clientInfoList, clientInfo)
		}
	}
	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "success",
		Data: clientInfoList,
	})
}

func (c *MsApi) SyncSlaveLevelStatus(ctx *gin.Context) {
	ms.Server.SyncClientLevelStatus()
	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "success",
		Data: nil,
	})
}
