package settings

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/wonderivan/logger"
)

var Conf = new(AppConf)

type AppConf struct {
	Mode       string `mapstructure:"mode"`
	Port       int    `mapstructure:"port"`
	*MysqlConf `mapstructure:"mysql"`
}

type MysqlConf struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Db           string `mapstructure:"db"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func InitConf() (err error) {
	viper.SetConfigFile("./conf/config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		logger.Error(errors.New("读取配置文件失败, " + err.Error()))
		return errors.New("读取配置文件失败, " + err.Error())
	}
	err = viper.Unmarshal(Conf)
	if err != nil {
		logger.Error(errors.New("配置反序列化失败, " + err.Error()))
		return errors.New("配置反序列化失败, " + err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logger.Error(errors.New("配置被修改, " + err.Error()))
		viper.Unmarshal(Conf)
	})
	return nil
}
