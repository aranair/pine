package config

import (
  "github.com/spf13/viper"
  "fmt"
  "log"
)

type Config struct {
  Coins []coin
}

type coin struct {
  Name string
  Cost float64
  Units float64
}

func LoadConfiguration(file string) Config {
  viper.SetConfigType("yaml")
  viper.SetConfigFile(file)

  if err := viper.ReadInConfig(); err != nil {
    log.Fatalf("Error reading config file, %s", err)
  }

  fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

  var C Config
  err := viper.Unmarshal(&C)
  if err != nil {
    panic(err)
  }

  // viper.WatchConfig()

  // viper.OnConfigChange(func(e fsnotify.Event) {
  //   fmt.Println("Config file changed:", e.Name)
  // })

  return C
}
