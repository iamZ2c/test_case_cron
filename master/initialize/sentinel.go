package initialize

import (
	"cron_tab_c/master/common"
	"cron_tab_c/master/global"
	"errors"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
)

func RateLimitClient() {
	var (
		err      error
		ruleList []*flow.Rule
	)
	global.SLogger.Infow("sentinel初始化配置")
	if err = sentinel.InitDefault(); err != nil {
		global.SLogger.Errorf("sentinel初始化配置失败:%v", err)
	}

	ruleList = make([]*flow.Rule, 0)

	for _, conf := range global.LimiterConfig {
		rule := flow.Rule{
			Resource:               conf.Resource,
			StatIntervalInMs:       conf.StatIntervalInMs,
			Threshold:              conf.Threshold,
			TokenCalculateStrategy: flow.Direct,
			// 匀速，不丢连接
			ControlBehavior: flow.Throttling,
		}
		ruleList = append(ruleList, &rule)
	}

	if _, err = flow.LoadRules(
		ruleList,
	); err != nil {
		global.SLogger.Errorf("sentinel加载配置失败:%v", err)
	}
}

func GetApiLimiter(name string) (rsc common.RateLimitServerConfig, err error) {

	for _, v := range global.LimiterConfig {
		if v.Resource == name {
			return v, nil
		}
	}
	return rsc, errors.New("未获取到对应配置信息")
}
