package static_getty

import (
	getty "github.com/apache/dubbo-getty"
	"github.com/huanglishi/gofly-mqttv5/cluster"
	"github.com/huanglishi/gofly-mqttv5/cluster/stat/colong"
	"github.com/huanglishi/gofly-mqttv5/cluster/stat/util"

	gxsync "github.com/dubbogo/gost/sync"
)

type server struct {
	name string
	s    getty.Server
}

func (this *server) Close() error {
	this.s.Close()
	return nil
}

var taskPool gxsync.GenericTaskPool

// RunClusterGettyServer 提供集群节点连接服务
func RunClusterGettyServer(name string, addr string, clusterInToPub colong.ClusterInToPub,
	clusterInToPubShare colong.ClusterInToPubShare, clusterInToPubSys colong.ClusterInToPubSys,
	shareTopicMapNode cluster.ShareTopicMapNode, taskPoolSize int) colong.NodeServerFace {
	util.SetLimit()

	//util.Profiling(*pprofPort)

	options := []getty.ServerOption{getty.WithLocalAddress(addr)}

	taskPool = gxsync.NewTaskPoolSimple(taskPoolSize)
	options = append(options, getty.WithServerTaskPool(taskPool))

	sev := getty.NewTCPServer(options...)

	go sev.RunEventLoop(newHelloServerSession(name, clusterInToPub, clusterInToPubShare,
		clusterInToPubSys, shareTopicMapNode))
	return &server{
		name: name,
		s:    sev,
	}
}

func newHelloServerSession(name string, clusterInToPub colong.ClusterInToPub,
	clusterInToPubShare colong.ClusterInToPubShare, clusterInToPubSys colong.ClusterInToPubSys,
	shareTopicMapNode cluster.ShareTopicMapNode) func(session getty.Session) error {
	return func(session getty.Session) error {
		err := InitialSessionServer(name, session, clusterInToPub, clusterInToPubShare,
			clusterInToPubSys, shareTopicMapNode)
		if err != nil {
			return err
		}
		return nil
	}
}
