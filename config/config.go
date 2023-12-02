package config

type Config struct {
	ServerAddress string
	DBSource      string
}

func LoadConfig() *Config {
	// 从环境变量或配置文件加载配置
	return &Config{
		ServerAddress: ":8080",
		DBSource:      "user:password@tcp(localhost:3306)/dbname",
	}
}
