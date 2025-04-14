package account

import "github.com/spf13/viper"

type Config struct {
	Port                string `mapstructure:"PORT"`
	DBDriver            string `mapstructure:"DB_DRIVER"`
	DBSourceDevelopment string `mapstructure:"DB_SOURCE_DEVELOPMENT"`
	DBSourceProduction  string `mapstructure:"DB_SOURCE_PRODUCTION"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile("env/account.env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
