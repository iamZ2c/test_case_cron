package common

type TotalConfig struct {
	EtcdConfig       EtcdConfig       `json:"etcd_config"`
	CronServerConfig CronServerConfig `json:"cron_server_config"`
}

type EtcdConfig struct {
	Addr string `json:"addr"`
	IP   string `json:"ip"`
}

type CronServerConfig struct {
	Host string `json:"host"`
	IP   string `json:"ip"`
}

type RateLimitServerConfig struct {
	Resource         string  `json:"resource"`
	StatIntervalInMs uint32  `json:"stat_interval_in_ms"`
	Threshold        float64 `json:"threshold"`
}

type RateLimitConfig []RateLimitServerConfig
