package config

type ServerConfig struct {
	ProjectName string `mapstructure:"project_name" json:"project_name"`
	Version     string `mapstructure:"version" json:"version"`
	Env         Env    `mapstructure:"env" json:"env"`
}

type Env struct {
	GO111MODULE string `mapstructure:"GO111MODULE" json:"GO111MODULE"`
	GOPROXY     string `mapstructure:"GOPROXY" json:"GOPROXY"`
}