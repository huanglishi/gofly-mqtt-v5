package authplus

import "github.com/huanglishi/gofly-mqttv5/logger"

type defaultAuth struct {
	i int
}

func defaultAuthPlusInit() {
	Register("", NewDefaultAuth()) //开启默认验证
	logger.Logger.Info("开启default进行增强认证")
}
func NewDefaultAuth() AuthPlus {
	return &defaultAuth{}
}
func (d2 *defaultAuth) Verify(authData []byte) (d []byte, continueAuth bool, err error) {
	if d2.i < 5 {
		d2.i++
		return []byte("continue"), true, nil
	} else {
		d2.i = 0
		return nil, false, nil
	}
}
