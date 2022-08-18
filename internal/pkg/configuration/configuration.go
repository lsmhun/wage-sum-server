package configuration

import (
	"log"
)

//viper "github.com/spf13/viper"

var configMap map[string]string

func init() {
	/*viper.SetConfigName("config")         // name of config file (without extension)
	//viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		log.Fatalf("fatal error config file: %w", err)
	}
	*/
	configMap = make(map[string]string)
	configMap["wagesum.db.host"] = "127.0.0.1"
	configMap["wagesum.db.port"] = "5432"
	configMap["wagesum.db.name"] = "wagesum"
	configMap["wagesum.http.service.port"] = "3000"
}

func GetConfigValue(key string) string {
	_, prs := configMap[key]
	if prs == false {
		log.Fatalf("Missing configuration: %s", key)
	}
	return configMap[key]
}
