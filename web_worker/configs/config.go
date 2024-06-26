package configs

import "github.com/spf13/viper"

type conf struct {
	KAFKAUrl   string `mapstructure:"KAFKA_URL"`
	KAFKATopic string `mapstructure:"KAFKA_TOPIC"`
	WebAppUrl  string `mapstructure:"WEB_APP_URL"`
}

func LoadConfig(paths []string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
