package app

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

var Config appConfig

type appConfig struct {
	MysqlDbDns string `mapstructure:"mysql_db_dns"`
	LogPath    string `mapstructure:"log_path"`
	Server     Server `mapstructure:"server"`
	Redis      Redis  `mapstructure:"redis"`
	Es         Es     `mapstructure:"es"`
	Jaeger     Jaeger `mapstructure:"jaeger"`
}

type Es struct {
	Dns          string `mapstructure:"dns"`
	EsGoodsIndex string `mapstructure:"es_goods_index"`
}

type Rabbit struct {
	Addr  string `mapstructure:"addr"`
	Vhost string `mapstructure:"vhost"`
}

type Server struct {
	Port string `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type Redis struct {
	HostName string `mapstructure:"hostname"`
	DB       int    `mapstructure:"database"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}
type Jaeger struct {
	Addr string `mapstructure:"addr"`
}

func LoadConfig(configPaths ...string) {
	v := viper.New()
	if os.Getenv("CHUANSHUI_ENV") == "prod" {
		v.SetConfigName("prod")
		v.SetConfigType("yaml")
	} else {
		v.SetConfigName("test")
		v.SetConfigType("yaml")
	}

	v.SetDefault("mysql_db_dns", "127.0.0.1")
	v.SetDefault("log_path", "./logs/runtime")

	v.SetDefault("redis.hostname", "127.0.0.1")
	v.SetDefault("redis.database", 1)
	v.SetDefault("redis.port", 6379)
	v.SetDefault("redis.password", "")

	v.SetDefault("jaeger.addr", "")

	v.SetDefault("server.port", 8080)
	v.SetDefault("server.name", "test")
	v.SetDefault("server.env", "test")

	v.SetDefault("rabbit.addr", "127.0.0.1:5672")
	v.SetDefault("rabbit.vhost", "test")

	v.SetDefault("es.dns", "")
	v.SetDefault("es.es_goods_index", "")

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		log.Panic(fmt.Errorf("config error failed to read the configuration file: %s", err))
	}
	if err := v.Unmarshal(&Config); err != nil {
		log.Panic("config error", err)
	}
}
