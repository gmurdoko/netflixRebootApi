package utils

import "github.com/spf13/viper"

//ViperGetEnv get data from environmet
func ViperGetEnv(key, defaultValue string) string {
	viper.AutomaticEnv()
	viper.SetConfigFile("./files/.env")
	viper.ReadInConfig()

	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}
