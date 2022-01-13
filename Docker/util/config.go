package util

import "github.com/spf13/viper"

type Config struct {
	AppPort string `mapstructure:"APP_PORT"`

	RuntimeSetup string `mapstructure:"RUNTIME_SETUP"`

	DBDriver string `mapstructure:"DB_DRIVER"`

	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {

	//adding app.env variables into struct()

	viper.AddConfigPath(path)

	viper.SetConfigName("app")

	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {

		print("Error detected!", err, "V", err.Error())

	}

	//save values into struct{} or map{} using umarshal

	viper.Unmarshal(&config)

	print("Params:", config.AppPort, config.RuntimeSetup)

	return config, err

}
