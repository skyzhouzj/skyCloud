package configs

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	SkyCloud   SkyCloud   `mapstructure:"SkyCloud"`
	DataSource DataSource `json:"DataSource"`
	Redis      Redis      `json:"Redis"`
}
type SkyCloud struct {
	ProjectName string `json:"ProjectName"`
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
func init() {
	viper.SetConfigName("dev_config.properties")
	viper.SetConfigType("properties")
	//默认取工作目录中的yaml
	viper.AddConfigPath("../")

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

func (*Config) Get() Config {
	return *config
}
