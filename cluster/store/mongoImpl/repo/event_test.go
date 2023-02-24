package mongorepo

import (
	"context"
	"fmt"
	"testing"

	"github.com/huanglishi/gofly-mqttv5/cluster/common"
	"github.com/huanglishi/gofly-mqttv5/config"
	"github.com/stretchr/testify/require"
)

// SI_CFG_PATH=F:/Go_pro/src/si-mqtt/config
func TestEvent(t *testing.T) {
	e := NewEventStore()
	ctx := context.Background()
	_ = e.Start(ctx, config.SIConfig{})
	require.NoError(t, e.PubEvent(ctx, common.NewEvent("12<", "123123", common.Connect)))
	p := common.NewEvent("12<", "123123", common.Sub)
	p.SubUnSub("tt", 0)
	require.NoError(t, e.PubEvent(ctx, p))
	fmt.Println(e.GetEventMaxTime(ctx))
	fmt.Println(e.GetEvent(ctx, 0))
}
