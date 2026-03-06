package config

type AppConfig struct {
	Mysql struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
}
type NacosConfig struct {
	Addr      string
	Port      int
	Namespace string
	DataID    string
	Group     string
	Servers   []NacosServer `mapstructure:"Servers"` // 多服务器配置
}

type NacosServer struct {
	Addr string
	Port int
}
