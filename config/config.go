package config

type Config struct {
  Coins coin
}

type coin struct {
	Symbol     string `json:"symbol"`
  Name     string `json:"name"`
	cost string `json:"cost"`
}

func LoadConfiguration(file string) Config {
    var config Config
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config
}
