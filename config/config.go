package config

import "github.com/spf13/viper"

type Config struct {
	EnvironmentVariables EnvironmentVariables `yaml:"EnvironmentVariables"`
	HTTPServer           HTTPServer           `yaml:"HTTPServer"`
	DataBase             DataBase             `yaml:"DataBase"`
	GRPCClient           GRPCClient           `yaml:"GRPCClient"`
	GRPCServer           GRPCServer           `yaml:"GRPCServer"`
	Logs                 Logs                 `yaml:"Logs"`
	Cors                 Cors                 `yaml:"Cors"`
}

type EnvironmentVariables struct {
	Environment string `yaml:"Environment"`
}

type HTTPServer struct {
	Addr string `yaml:"Addr"`
	Port string `yaml:"Port"`
}

type DataBase struct {
	ConnectionString   string `yaml:"ConnectionString"`
	ConnectionPostgres string `yaml:"ConnectionPostgres"`
	Name               string `yaml:"Name"`
}

type GRPCClient struct {
	Services map[string]string `yaml:"Services"`
}

type GRPCServer struct {
	Type string
	Addr string
}

type Logs struct {
	Path       string `yaml:"Path"`
	Level      string `yaml:"Level"`
	MaxAge     int    `yaml:"MaxAge"`
	MaxBackups int    `yaml:"MaxBackups"`
}

type Cors struct {
	AllowedOrigins []string `yaml:"AllowedOrigins"`
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
