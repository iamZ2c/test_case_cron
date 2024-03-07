package initialize

import (
	"cron_tab_c/master/global"
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func NacosClient() (err error) {
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	constantConfig := constant.ClientConfig{
		NamespaceId:         "d4d8c2d7-6247-41e2-8723-2e94395c0855", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           3000,
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		CacheDir:            "cache",
		LogLevel:            "debug",
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  constantConfig,
	})
	if err != nil {
		global.SLogger.Errorf("dddd")
		return err
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "cron_server_limiter",
		Group:  "test",
	})

	if err = json.Unmarshal([]byte(content), &global.LimiterConfig); err != nil {
		global.SLogger.Errorf("解析nacos,json文件失败")
		return err
	}

	content, err = configClient.GetConfig(vo.ConfigParam{
		DataId: "server_config",
		Group:  "test",
	})
	if err = json.Unmarshal([]byte(content), &global.TotalConfig); err != nil {
		global.SLogger.Errorf("解析nacos,json文件失败")
		return err
	}

	if err = configClient.ListenConfig(
		vo.ConfigParam{
			DataId: "cron_server_limiter",
			Group:  "test",
			OnChange: func(namespace, group, dataId, data string) {
				content, err := configClient.GetConfig(vo.ConfigParam{
					DataId: "cron_server_limiter",
					Group:  "test",
				})
				if err = json.Unmarshal([]byte(content), &global.LimiterConfig); err != nil {
					global.SLogger.Errorf("解析nacos,json文件失败")
				}
				global.SLogger.Infow("更新cron_server_limiter")
			},
		},
	); err != nil {
		return err
	}
	return nil
}
