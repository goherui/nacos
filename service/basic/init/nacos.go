package init

import (
	"day6/service/basic/config"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

// getProjectRoot 获取项目根目录
func getProjectRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		dir := filepath.Dir(filename)
		for i := 0; i < 3; i++ {
			dir = filepath.Dir(dir)
		}
		return dir
	}
	wd, _ := os.Getwd()
	return wd
}

func InitNacos() {
	var err error
	configPath := filepath.Join(getProjectRoot(), "nacos.yaml")
	viper.SetConfigFile(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var NacosConfig config.NacosConfig
	err = viper.UnmarshalKey("Nacos", &NacosConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("nacos配置成功", NacosConfig)

	// Nacos服务器地址
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: NacosConfig.Addr,
			Port:   uint64(NacosConfig.Port),
		},
	}

	// 构建Nacos客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: false,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建Nacos配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	configContent, err := configClient.GetConfig(vo.ConfigParam{
		DataId: NacosConfig.DataID,
		Group:  NacosConfig.Group,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("从Nacos读取到的配置内容:\n", configContent)
	config.GlobalConfig = &config.AppConfig{}
	nacosViper := viper.New()
	nacosViper.SetConfigType("yaml")
	err = nacosViper.ReadConfig(strings.NewReader(configContent))
	if err != nil {
		panic(fmt.Errorf("加载Nacos配置内容到viper失败: %w", err))
	}
	err = nacosViper.Unmarshal(config.GlobalConfig)
	if err != nil {
		panic(fmt.Errorf("将Nacos配置解析到GlobalConfig失败: %w", err))
	}
	fmt.Println("配置加载成功", config.GlobalConfig)
}
