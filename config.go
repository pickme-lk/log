package log

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type LogConfig struct {
	Level         string `yaml:"level" json:"level"`
	RemoteLogging bool   `yaml:"remote_logging" json:"remote_logging"`
	FilePath      bool   `yaml:"file_path_enabled" json:"file_path_enabled"`
	Colors        bool   `yaml:"colors" json:"colors"`
}

var Config LogConfig

func init() {
	setDefaults()
	loadConfig()

	if os.Getenv(`LOG_LEVEL`) != `` {
		Config.Level = os.Getenv(`LOG_LEVEL`)
	}

	if os.Getenv(`LOG_FILE_PATH`) != `` && os.Getenv(`LOG_FILE_PATH`) == `1` {
		Config.FilePath = true
	}

}

func setDefaults() {
	Config = LogConfig{}
	Config.Level = `INFO`
	Config.Colors = true
	Config.FilePath = true
}

//loadConfig Load logging configurations
func loadConfig() {

	file, err := ioutil.ReadFile(`config/logger.yaml`)
	if err != nil {
		log.Println(`go-util/log: Cannot open config file `, err)
		return
	}

	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		log.Fatalln(`go-util/log: Cannot decode config file `, err)
	}
}
