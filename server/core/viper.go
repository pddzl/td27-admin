package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"server/core/internal"
	"server/global"
)

func Viper() *viper.Viper {
	config := internal.ConfigFile

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.TD27_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.TD27_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
