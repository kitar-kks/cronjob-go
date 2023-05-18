package configuration

import (
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {
	Schedulers []Scheduler `mapstructure:"schedulers"`
}

type Scheduler struct {
	Job  string `mapstructure:"job"`
	Cron string `mapstructure:"cron"`
	Task struct {
		URL    string `mapstructure:"url"`
		Method string `mapstructure:"method"`
		Body   string `mapstructure:"body"`
		Header string `mapstructure:"header"`
	} `mapstructure:"task"`
}

var Config Configuration

func Load(env string) {
	viper.SetConfigName(env)
	viper.AddConfigPath("configuration")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	_ = viper.Unmarshal(&Config)
}
