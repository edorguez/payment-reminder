package account

import "github.com/spf13/viper"

type Config struct {
	Environment           string `mapstructure:"ENVIRONMENT"`
	Account_Svc_Port      string `mapstructure:"ACCOUNT_SVC_PORT"`
	DB_Source_Development string `mapstructure:"DB_SOURCE_DEVELOPMENT"`
	DB_Source_Production  string `mapstructure:"DB_SOURCE_PRODUCTION"`
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
