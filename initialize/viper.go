package initialize

import (
	"flag"
	"fmt"
	"gin-practice/config"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var vip *viper.Viper

func init() {
	var cfg string
	flag.StringVar(&cfg, "c", "", "")
	flag.Parse()
	if cfg == "" {
		fmt.Println("cfg null!")
		if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
			switch gin.Mode() {
			case gin.DebugMode:
				cfg = ConfigDebugFile
			case gin.ReleaseMode:
				cfg = ConfigReleaseFile
			case gin.TestMode:
				cfg = ConfigTestFile
			}
		} else {
			cfg = configEnv
		}
	}
	fmt.Println("cfg value: ", cfg)
	vip = viper.New()
	vip.SetConfigFile(cfg)
	err := vip.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if err = vip.Unmarshal(&config.SELF_CONFIG); err != nil {
		panic(err)
	}
}

func GetConfig() *viper.Viper {
	return vip
}
