package configig

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// В конфиге сохраняем адрес и базу данных
type Config struct {
	Addr  string `yaml:"port"`
	DbUrl string `yaml:"db_url"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadConfig(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open config file :%v", err)
	}
	defer file.Close()
	read, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed reading config file :%v", err)
	}
	err = yaml.Unmarshal(read, c)
	if err != nil {
		return fmt.Errorf("yaml unmarshalling error :%v", err)
	}

	return nil
}
