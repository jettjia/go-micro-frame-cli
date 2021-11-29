package constant

const (
	VERSION     = "v1.2.5"
	PROJECTNAME = "go-micro-frame-cli"

	GO111MODULE = "on"
	GOPROXY = "https://goproxy.cn"
	GOVERSION = "1.16.7"

	DockerVersion = "19.03.*"

	MysqlVersion = "5.7"
	MysqlName = "mysql-jett"

	RedisVersion = "6.2"
	RedisName = "redis-jett"

	RabbitmqVersion = "3.7.7-management"
	RabbitmqName = "rabbit-jett"

	NacosVersion = "latest"
	NacosName = "nacos-jett"

	JaegerVersion = "latest"
	JaegerName = "jaeger-jett"

	ConsulVersion = "latest"
	ConsulName = "consul-jett"

	// install kong/konga
	//// konga依赖的 postgres数据库信息
	PostgresName = "kong-database-jett"
	PostgresDBUser = "kong"
	PostgresDB = "kong"
	PostgresPwd = "kong"
	//// kong/konga
	KongVersion = "2.5.0"
	KongaName = "konga-jett"

	EsVersion = "7.7.1"
	EsName = "es-jett"
)
