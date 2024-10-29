package config

import "github.com/spf13/viper"

type Config struct {
	EnvironmentVariables EnvironmentVariables `yaml:"EnvironmentVariables"`
	ConnectionStrings    ConnectionStrings    `yaml:"ConnectionStrings"`
	HTTPServer           HTTPServer           `yaml:"HTTPServer"`
	Logs                 Logs                 `yaml:"Logs"`
}

type EnvironmentVariables struct {
	Environment string `yaml:"Environment"`
}

type ConnectionStrings struct {
	ServiceDb string `yaml:"ServiceDb"`
}

type HTTPServer struct {
	Addr string `yaml:"Addr"`
	Port string `yaml:"Port"`
}

type Logs struct {
	Path       string `yaml:"Path"`
	Level      string `yaml:"Level"`
	MaxAge     int    `yaml:"MaxAge"`
	MaxBackups int    `yaml:"MaxBackups"`
}

func ReadConfig(cfgName, cfgType, cfgPath string) (*Config, error) {
	var cfg Config

	viper.SetConfigName(cfgName)
	viper.SetConfigType(cfgType)
	viper.AddConfigPath(cfgPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
