package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ProviderSet = wire.NewSet(NewConf)



//func NewLoggerHelper(logger *log.Logger) *log.Helper  {
//	return log.NewHelper(*logger)
//}

type Conf struct {
	Conf *viper.Viper
}

func NewConf() *Conf {
	viper := viper.New()
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./configs")  // path to look for the config file in
	//vip.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")           // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                     // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	return &Conf{
		Conf: viper,
	}
}
