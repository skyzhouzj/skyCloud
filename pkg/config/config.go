package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type SkyCloud struct {
	ProjectName             string `json:"ProjectName"`
	ProjectVersion          string `json:"ProjectVersion"`
	ProjectDomain           string `json:"ProjectDomain"`
	ProjectPort             string `json:"ProjectPort"`
	ProjectAccessLogFile    string `json:"ProjectAccessLogFile"`
	HeaderLoginToken        string `json:"HeaderLoginToken"`
	HeaderSignToken         string `json:"HeaderSignToken"`
	HeaderSignTokenDate     string `json:"HeaderSignTokenDate"`
	RedisKeyPrefixRequestID string `json:"RedisKeyPrefixRequestID"`
	RedisKeyPrefixLoginUser string `json:"RedisKeyPrefixLoginUser"`
	RedisKeyPrefixSignature string `json:"RedisKeyPrefixSignature"`
}
type DataSource struct {
	DbType       string `json:"dbType"`
	Path         string `json:"path"`
	Config       string `json:"config"`
	DbName       string `json:"db-name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	MaxIdleConns int    `json:"max-idle-conns"`
	MaxOpenConns int    `json:"max-open-conns"`
	LogMode      bool   `json:"log-mode"`
}
type Redis struct {
	Db       int    `json:"db"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

//config加载
func Init(config interface{}, path string, envStr string) {
	viper.SetConfigName(envStr + "_config.properties")
	viper.SetConfigType("properties")
	//默认取工作目录中的yaml
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
	//监控配置文件是否修改
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}
