package ms

import (
	"context"
	pb "dst-admin-go/proto/messaging"
	"dst-admin-go/service"
	"dst-admin-go/utils/dstConfigUtils"
	"dst-admin-go/utils/shellUtils"
	"dst-admin-go/vo"
	"dst-admin-go/vo/level"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"runtime"
	"strings"
	"time"
)

func StartClient(target, name string) {
	for {
		if err := runClient(target, name); err != nil {
			log.Printf("Client encountered an error: %v", err)
			time.Sleep(1 * time.Minute) // 重试连接之前等待一段时间
		}
	}
}

func runClient(target string, name string) error {
	//target := "localhost:50051"
	//target := "139.159.184.218:50051"

	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return err
	}
	defer conn.Close()

	c := pb.NewMessagingServiceClient(conn)

	stream, err := c.Connect(context.Background())
	if err != nil {
		return err
	}

	// 发送初始消息，包含客户端ID
	if err := stream.Send(&pb.Message{ClientId: name, Content: "hi", Type: pb.MessageType_PRIVATE}); err != nil {
		return err
	}
	log.Println("[connect success]")
	stop := make(chan bool, 0)
	// 监听来自服务器的消息
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				log.Println("Stream closed by server")
				stop <- true
				return
			}
			if err != nil {
				log.Printf("Failed to receive message: %v", err)
				stop <- true
				return
			}
			if in.FnName == SYNC_LEVEL {
				syncModoverrides(in.Content)
			} else if in.FnName == SYNC_START {
				syncStartLevel(in.Content)
			} else if in.FnName == SERVER_SYNC_STATUS {
				levelStatus := syncLevelStatus()
				bytes, err2 := json.Marshal(levelStatus)
				if err2 != nil {
					log.Println(err2)
				} else {
					if err := stream.Send(&pb.Message{ClientId: name, Content: string(bytes), Type: pb.MessageType_PRIVATE, FnName: CLIENT_STATUS}); err != nil {
						log.Println(err)
					}
				}
			}
			log.Printf("Received message of type %s from server: %s", in.Type, in.Content)
		}
	}()

	<-stop
	return errors.New("退出")
}

var gameService service.GameService
var homeServer service.HomeService
var gameLevel2Service service.GameLevel2Service

func syncModoverrides(data string) {
	// cluster_ini 要修改
	// cluster_token 要修改
	syncLevelData := &SyncLevelData{}
	err := json.Unmarshal([]byte(data), &syncLevelData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	clusterName := dstConfigUtils.GetDstConfig().Cluster
	clusterIni := homeServer.GetClusterIni(clusterName)
	clusterIni.BindIp = "127.0.0.1"
	clusterIni.MasterIp = syncLevelData.MasterIp
	clusterIni.MasterPort = syncLevelData.MasterPort

	homeServer.SaveClusterIni(clusterName, clusterIni)
	homeServer.SaveClusterToken(clusterName, syncLevelData.ClusterToken)

	// 模组
	levelList := gameLevel2Service.GetLevelList(clusterName)
	for i := range levelList {
		levelList[i].Modoverrides = syncLevelData.Modoverrides
	}

	err = gameLevel2Service.UpdateLevels(clusterName, levelList)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}

type LevelInfo struct {
	Ps                *vo.DstPsVo      `json:"Ps"`
	Status            bool             `json:"status"`
	LevelName         string           `json:"levelName"`
	IsMaster          bool             `json:"is_master"`
	Uuid              string           `json:"uuid"`
	Leveldataoverride string           `json:"leveldataoverride"`
	Modoverrides      string           `json:"modoverrides"`
	ServerIni         *level.ServerIni `json:"server_ini"`
}

func syncLevelStatus() []LevelInfo {

	clusterName := dstConfigUtils.GetDstConfig().Cluster

	levelList := gameLevel2Service.GetLevelList(clusterName)
	length := len(levelList)
	result := make([]LevelInfo, length)

	if runtime.GOOS == "windows" {
		for i := range levelList {
			world := levelList[i]
			ps := gameService.PsAuxSpecified(clusterName, world.Uuid)
			status := gameService.GetLevelStatus(clusterName, world.Uuid)
			result[i] = LevelInfo{
				Ps:                ps,
				Status:            status,
				LevelName:         world.LevelName,
				IsMaster:          world.IsMaster,
				Uuid:              world.Uuid,
				Leveldataoverride: world.Leveldataoverride,
				Modoverrides:      world.Modoverrides,
				ServerIni:         world.ServerIni,
			}
		}
		return result
	}

	cmd := "ps -aux | grep -v grep | grep -v tail | grep -v SCREEN | grep " + clusterName + " |awk '{print $3, $4, $5, $6,$16}'"
	info, err := shellUtils.Shell(cmd)
	if err != nil {
		log.Println(cmd + " error: " + err.Error())
	} else {
		lines := strings.Split(info, "\n")
		for lineIndex := range lines {
			dstPsVo := vo.NewDstPsVo()
			arr := strings.Split(lines[lineIndex], " ")
			if len(arr) > 4 {
				dstPsVo.CpuUage = strings.Replace(arr[0], "\n", "", -1)
				dstPsVo.MemUage = strings.Replace(arr[1], "\n", "", -1)
				dstPsVo.VSZ = strings.Replace(arr[2], "\n", "", -1)
				dstPsVo.RSS = strings.Replace(arr[3], "\n", "", -1)
				for i := range result {
					levelName := result[i].Uuid
					if strings.Contains(arr[4], levelName) {
						result[i].Ps = dstPsVo
						result[i].Status = true
					}
				}
			}
		}
	}

	return result
}

func syncStartLevel(start string) {
	clusterName := dstConfigUtils.GetDstConfig().Cluster
	if start == "start" {
		gameService.StartGame(clusterName)
	}
	if start == "stop" {
		gameService.StopGame(clusterName)
	}
	syncLevelStatus()
}
