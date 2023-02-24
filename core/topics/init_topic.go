package topics

import (
	"github.com/huanglishi/gofly-mqttv5/core/topics/share"
	"github.com/huanglishi/gofly-mqttv5/core/topics/sys"
)

func TopicInit(topicPro string) {
	switch topicPro {
	default:
		sys.SysTopicInit()
		share.ShareTopicInit()
		memTopicInit() // 这个顺序必须在前两个后面
	}
	//redisTopicInit()
}
